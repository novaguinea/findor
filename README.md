#Findor API Documentation

##Setup Go on your laptop/pc

1. Download golang in this [link](https://go.dev/dl/)
2. Install golang (usual installation)
3. Make sure golang path added to your advance system settings (system variables path)
4. Check whether your golang successfully installed by typing this command on command line/terminal
`go version`

##Setup database for Findor

1. Open XAMPP Control Panel and start MySQL (start apache if you open it from phpMyAdmin)
2. Create database 'Findor'
3. Make sure findor database created!

##Setup Findor project

1. Clone this project 
`git clone https://github.com/novaguinea/findor.git`

2. Open your terminal and direct to findor's folder
3. Just to make sure all dependencies are installed run this in your command line (CLI)
`go mod tidy`

4. Open your XAMPP Control Panel and start MySQL
5. Run findor project by run this command in CLI/terminal
`go run main.go`
