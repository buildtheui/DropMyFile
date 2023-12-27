# DropMyFile

Share files from and to the device running DropMyFile using mobile or desktop devices

## Why?

Ever had to send files and ended up uploading them to cloud storage like drive or using email and messaging apps like whatsapp or telegram? It can be a hassle. With DropMyFile, you can do it faster by using the same Wi-Fi connection your devices are on (LAN Network) – no need to open a bunch of apps!

## How?

If your devices are on the same network, just fire up DropMyFile on your computer, scan the QR code, and upload your files using the opened webapp. You can also stash your files in DropMyFile's transfer folder and download them on another device using the rendered app from the QR code.

## What You Need

1. Make sure all your devices are on the same Wi-Fi network (LAN network).

## How to use it

1. **Download:** Choose the appropriate file for your platform from the [releases section](https://github.com/buildtheui/DropMyFile/releases).

2. **Run the File:** Execute the downloaded file through the terminal.

3. **Connect Devices:** Scan the provided QR code with another device or enter the displayed link.

4. **Upload Files:** Easily transfer files by uploading them. By default, these files will be stored in a folder named `transferedFiles` on the desktop folder.

## Customize

You can customize security, port, and the folder from enviroment variables or flags. take into account that environment variables will take prescedence from flags.

| Flags            | Short | Env Variable        | Default                                                      | Description                                                                                                                                                                                                |
| ---------------- | ----- | ------------------- | ------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| --folder-path    | -f    | DMF_TRANSFER_FOLDER | /Users/YourUsername/Desktop or C:\Users\YourUsername\Desktop | Folder where all transferred files are uploaded or downloaded from                                                                                                                                         |
| --port           | -p    | DMF_PORT            | 3000                                                         | Port running all file transfers                                                                                                                                                                            |
| --session-length | -s    | DMF_SESSION_LENGTH  | 6                                                            | A random string generated when DropMyFile is executed and appended to the QR link. Devices without this session string cannot upload or download files. This can be deactivated by setting the value to 0. |
