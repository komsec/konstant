* User experience
* Admin
    * AD
* Asset import
    * Network - 1000 linux systems
        * Each and every servers to be imported
            * Using CSV - name, ip address
        * Client side
            * Agent must be installed/managed
                * Agent will ask for server details and it will automatically register to server
                * No need of import
            * Agentless is there
                * Using SSH - using certificates
                    * Cyber arc through authentication
                * Windows wmi is used for agentless
            * Agent based
                * MTLS based authentication/security
* java interface
* Customizations
    * Config changes
        * Guest user disabled
        * SNMP community strings
        * AD 
    * Symantec endpoint protection
        * There are 4 services
    * Everything is written as regex
        * Maybe regex validations right there
* All are gui based
* Less abstractions
    * Systemd/init etc are separate functions

* Pros
    * 
* Cons
    * 
* Wishlist
    * APIs maybe developed
    * Asset auto discovery
    * Integration apis to be needed
        * for third party integration

* Agent based Pros
    * Faster
    * No import needed
* Agentless


* Basic Features for MVP
    * Asset import
        * inputs - hostname, ip address
        * Auto discovery - os type, kernel etc needed to be discovered
    * Custom schedule for compliance check etc weekly
    * Standard creation/updates
    * 
    * Report generation
        - csv, yaml, 
        - Push to servicenow, s3
    * Exceptions
        - some app would need Exceptions

