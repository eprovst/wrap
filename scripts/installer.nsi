; This is a temporary solution, until WiX support is better on Linux.

Unicode True

!define Name "Wrap"
Name "${Name}"
Outfile "../dist/${Name}_Win64_nightly.exe"
RequestExecutionLevel admin ;Require admin rights on NT6+ (When UAC is turned on)
InstallDir "$ProgramFiles64\${Name}"
!define MUI_ICON "../assets/wrap/wrap.ico"

!include LogicLib.nsh
!include MUI.nsh
!include x64.nsh

Function .onInit
SetShellVarContext all
UserInfo::GetAccountType
pop $0
${If} $0 != "admin" ;Require admin rights on NT4+
    MessageBox mb_iconstop "Administrator rights required!"
    SetErrorLevel 740 ;ERROR_ELEVATION_REQUIRED
    Quit
${EndIf}

${IfNot} ${RunningX64}
    MessageBox mb_iconstop "This version of Wrap can only be installed on 64bit Windows!"
    SetErrorLevel 2
    Quit
${EndIf}
FunctionEnd

!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH
!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES
!insertmacro MUI_LANGUAGE "English"

; The following functions have been provided by smartmontools.org:
; Registry Entry for environment (NT4,2000,XP)
; All users:
!define Environ 'HKLM "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"'

; AddToPath - Appends dir to PATH
;   (does not work on Win9x/ME)
;
; Usage:
;   Push "dir"
;   Call AddToExecPath

Function AddToExecPath
  Exch $0
  Push $1
  Push $2
  Push $3
  Push $4

  ; NSIS ReadRegStr returns empty string on string overflow
  ; Native calls are used here to check actual length of PATH

  ; $4 = RegOpenKey(HKEY_LOCAL_MACHINE, ".....\Environment", &$3)
  System::Call "advapi32::RegOpenKey(i 0x80000002, t'SYSTEM\CurrentControlSet\Control\Session Manager\Environment', *i.r3) i.r4"
  IntCmp $4 0 0 done done
  ; $4 = RegQueryValueEx($3, "PATH", (DWORD*)0, (DWORD*)0, &$1, ($2=NSIS_MAX_STRLEN, &$2))
  ; RegCloseKey($3)
  System::Call "advapi32::RegQueryValueEx(i $3, t'PATH', i 0, i 0, t.r1, *i ${NSIS_MAX_STRLEN} r2) i.r4"
  System::Call "advapi32::RegCloseKey(i $3)"

  IntCmp $4 234 0 +4 +4 ; $4 == ERROR_MORE_DATA
  DetailPrint "AddToPath: original length $2 > ${NSIS_MAX_STRLEN}"
  MessageBox MB_OK "PATH not updated, original length $2 > ${NSIS_MAX_STRLEN}"
  Goto done

  IntCmp $4 0 +5 ; $4 != NO_ERROR
  IntCmp $4 2 +3 ; $4 != ERROR_FILE_NOT_FOUND
  DetailPrint "AddToPath: unexpected error code $4"
  Goto done
  StrCpy $1 ""

  ; Check if already in PATH
  Push "$1;"
  Push "$0;"
  Call StrStr
  Pop $2
  StrCmp $2 "" 0 done
  Push "$1;"
  Push "$0\;"
  Call StrStr
  Pop $2
  StrCmp $2 "" 0 done

  ; Prevent NSIS string overflow
  StrLen $2 $0
  StrLen $3 $1
  IntOp $2 $2 + $3
  IntOp $2 $2 + 2 ; $2 = strlen(dir) + strlen(PATH) + sizeof(";")
  IntCmp $2 ${NSIS_MAX_STRLEN} +4 +4 0
  DetailPrint "AddToPath: new length $2 > ${NSIS_MAX_STRLEN}"
  MessageBox MB_OK "PATH not updated, new length $2 > ${NSIS_MAX_STRLEN}."
  Goto done

  ; Append dir to PATH
  DetailPrint "Add to PATH: $0"
  StrCpy $2 $1 1 -1
  StrCmp $2 ";" 0 +2
  StrCpy $1 $1 -1 ; remove trailing ';'
  StrCmp $1 "" +2   ; no leading ';'
  StrCpy $0 "$1;$0"
  WriteRegExpandStr ${Environ} "PATH" $0
  SendMessage ${HWND_BROADCAST} ${WM_WININICHANGE} 0 "STR:Environment" /TIMEOUT=5000

done:
  Pop $4
  Pop $3
  Pop $2
  Pop $1
  Pop $0
FunctionEnd


; RemoveFromPath - Removes dir from PATH
;
; Usage:
;   Push "dir"
;   Call RemoveFromExecPath

Function un.RemoveFromExecPath
  Exch $0
  Push $1
  Push $2
  Push $3
  Push $4
  Push $5
  Push $6

  ReadRegStr $1 ${Environ} "PATH"
  StrCpy $5 $1 1 -1
  StrCmp $5 ";" +2
  StrCpy $1 "$1;" ; ensure trailing ';'
  Push $1
  Push "$0;"
  Call un.StrStr
  Pop $2 ; pos of our dir
  StrCmp $2 "" done

  DetailPrint "Remove from PATH: $0"
  StrLen $3 "$0;"
  StrLen $4 $2
  StrCpy $5 $1 -$4 ; $5 is now the part before the path to remove
  StrCpy $6 $2 "" $3 ; $6 is now the part after the path to remove
  StrCpy $3 "$5$6"
  StrCpy $5 $3 1 -1
  StrCmp $5 ";" 0 +2
  StrCpy $3 $3 -1 ; remove trailing ';'
  WriteRegExpandStr ${Environ} "PATH" $3
  SendMessage ${HWND_BROADCAST} ${WM_WININICHANGE} 0 "STR:Environment" /TIMEOUT=5000

done:
  Pop $6
  Pop $5
  Pop $4
  Pop $3
  Pop $2
  Pop $1
  Pop $0
FunctionEnd

; StrStr - find substring in a string
;
; Usage:
;   Push "this is some string"
;   Push "some"
;   Call StrStr
;   Pop $0 ; "some string"

!macro StrStr un
Function ${un}StrStr
  Exch $R1 ; $R1=substring, stack=[old$R1,string,...]
  Exch     ;                stack=[string,old$R1,...]
  Exch $R2 ; $R2=string,    stack=[old$R2,old$R1,...]
  Push $R3
  Push $R4
  Push $R5
  StrLen $R3 $R1
  StrCpy $R4 0
  ; $R1=substring, $R2=string, $R3=strlen(substring)
  ; $R4=count, $R5=tmp
  loop:
    StrCpy $R5 $R2 $R3 $R4
    StrCmp $R5 $R1 done
    StrCmp $R5 "" done
    IntOp $R4 $R4 + 1
    Goto loop
done:
  StrCpy $R1 $R2 "" $R4
  Pop $R5
  Pop $R4
  Pop $R3
  Pop $R2
  Exch $R1 ; $R1=old$R1, stack=[result,...]
FunctionEnd
!macroend
!insertmacro StrStr ""
!insertmacro StrStr "un."

Section
SetOutPath "$INSTDIR"

WriteUninstaller "$INSTDIR\Uninstall.exe"
WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Wrap"   "DisplayName" "${Name}"
WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Wrap"   "UninstallString" "$INSTDIR\Uninstall.exe"

; File types and their icons
WriteRegStr HKCR ".fountain" "" "Wrap.Fountain"
WriteRegStr HKCR "Wrap.Fountain" "" "Fountain"
WriteRegStr HKCR "Wrap.Fountain\DefaultIcon" "" "$INSTDIR\fountain.ico"
File /oname=fountain.ico ../assets/filetypes/fountain/fountain.ico

WriteRegStr HKCR ".wrap" "" "Wrap.Wrap"
WriteRegStr HKCR "Wrap.Wrap" "" "Wrap"
WriteRegStr HKCR "Wrap.Wrap\DefaultIcon" "" "$INSTDIR\wrap.ico"
File /oname=wrap.ico ../assets/filetypes/wrap/wrap.ico

; Executable
File /oname=wrap.exe ../build/windows/wrap.exe

Push $INSTDIR
Call AddToExecPath
SectionEnd


Section "uninstall"

; Executable
Delete "$INSTDIR\wrap.exe"

; File type icons
DeleteRegKey HKCR ".fountain"
DeleteRegKey HKCR "Wrap.Fountain"
Delete "$INSTDIR\fountain.ico"

DeleteRegKey HKCR ".wrap"
DeleteRegKey HKCR "Wrap.Wrap"
Delete "$INSTDIR\wrap.ico"

DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\Wrap"
Push $INSTDIR
Call un.RemoveFromExecPath

Delete "$INSTDIR\Uninstall.exe"
RMDir "$INSTDIR"
SectionEnd
