# Custom Portable of WaykNow

## Requirements
- [GoLang](https://golang.org/dl/)
- [Statik CLI](https://github.com/rakyll/statik) (embeds files)
- [Latest WaykNow EXE](https://wayk.devolutions.net/wayk-now/home/thankyou/waykbin)
- [Custom Made WaykNow JSON Config File](https://helpwayk.devolutions.net/kb_configcommandline.html)

## Go Requirements
 - [configdir](https://github.com/kirsle/configdir) `go get github.com/kirsle/configdir`
 - [statik](https://github.com/rakyll/statik) `go get github.com/rakyll/statik`

## Instructions
 1. clone repo
 2. place the WaykNow EXE inside the `public` folder named as `wayknow.exe`
 3. edit the `WaykNow.cfg` which is located inside the `public` folder with your custom configs
 4. run command prompt or powershell at the clones repo folder
 5. run `statik` (this will build the embedded files required)
 6. **RUN STEP 7 OR 8, NOT BOTH!**
 7. run `go build -ldflags="-H windowsgui" -o portable.exe` (this builds the app **WITHOUT A GUI**)
 8. run `go build -o portable.exe` (this builds the app **WITH A BLACK GUI WINDOW FOR PROGRESS**)

