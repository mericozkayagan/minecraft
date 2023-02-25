# Minecraft CLI
This cli launchs an EC2 with a Minecraft server configured. It works with the credentials from .aws/credentials. At the moment it only supports the _eu-west-1_ and requires to subscription to https://aws.amazon.com/marketplace/pp/prodview-kt5vwxevivova

### The commands are:
* **minecraft install**
   - This command creates a security group with necessary ports and launchs the instance.
     **Inputs:**
    - Instance type >> t2.medium, t3.small etc
* **minecraft destroy**
    - This commands destroys the instance and the related resources
* **minecraft reboot**
    - Reboots the instance if anything goes wrong
* **minecraft start**
    - Starts the instance
* **minecraft stop**
    - Stops the instance
