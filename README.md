# Project Setup and Running a Cron Job

This README provides detailed instructions for setting up project and scheduling a cron job to automate running the project.

## Project Setup

### 1. Prerequisites

Ensure you have the following installed on your system:

Go (Golang) installed. You can download it from https://go.dev/dl/.

A terminal or command-line interface.

### 2. Clone the Repository

Clone the project repository to your local system:

git clone `https://github.com/ayushgupta-lang/updateNumber`


###  3. Navigate to the Project Directory

Change to the project directory where the main.go file is located:

cd ~/update-number-project

### 4. Run the Project Manually (Optional)

To ensure the project runs correctly, execute it manually:

`go run main.go`

If the script runs successfully, you can proceed to schedule it as a cron job.

## Setting Up a Cron Job

### 1. Open the Crontab Editor

To open the crontab editor for your user:

`crontab -e`

If it’s your first time using crontab, it will prompt you to select a text editor (choose your preferred one, such as nano or vim).

### 2. Add the Cron Job

In the crontab editor, add the following line to schedule the script to run daily at midnight:

`0 0 * * * cd ~/update-number-project && go run main.go`

### 3. Save and Exit

In nano:

Press Ctrl + O to save, then Enter to confirm.

Press Ctrl + X to exit.

In vim/vi:

Press Esc, then type :wq and press Enter to save and exit.

### 4. Verify the Cron Job

To confirm that the cron job has been added successfully, run:

`crontab -l`

You should see the line you just added.

## Removing the Cron Job

### 1. Open the Crontab Editor

To remove the cron job, open the crontab editor:

`crontab -e`

### 2. Delete the Line

Locate the line you added earlier:

`0 0 * * * cd ~/update-number-project && go run main.go`

Delete the line using your editor.

### 3. Save and Exit

Follow the same steps to save and exit as described above.

### 4. Verify Deletion

To ensure the cron job was removed, run:

`crontab -l`

If the job was deleted successfully, the output will not include the removed line.

Additional Commands

List Existing Cron Jobs

To list all cron jobs for the current user:

`crontab -l`

Edit the Crontab

To edit the crontab file:

`crontab -e`

Remove All Cron Jobs

To delete all cron jobs for the current user:

`crontab -r`

Run the Script Immediately (For Testing)

If you want to test the script without scheduling it:

`cd ~/update-number-project && go run main.go`