Controllers found: 1
----------------------------------------------------------------------
Controller information
----------------------------------------------------------------------
   Controller Status                        : Optimal
   Controller Mode                          : Mixed
   Channel description                      : SCSI
   Controller Model                         : Adaptec SmartIOC 8i
   Controller Serial Number                 : Unknown
   Controller World Wide Name               : 50123456789ABC00
   Physical Slot                            : 10
   Temperature                              : 48 C/ 118 F (Normal)
   Power Consumption                        : Not Applicable
   Power Mode                               : Unknown
   Power Mode Operational                   : Unknown
   Survival Mode                            : Unknown
   Host bus type                            : PCIe 3.0
   Host bus speed                           : 7880 MBps
   Host bus link width                      : 8 bit(s)/link(s)
   PCI Address (Bus:Device:Function)        : 0:1:0:0
   I2C Address                              : 0xDE
   I2C Clock Speed                          : 400 KHZ
   I2C Clock Stretching                     : Enabled
   Number of Ports                          : 2
   Internal Port Count                      : 2
   External Port Count                      : 0
   Defunct disk drive count                 : 0
   Logical devices/Failed/Degraded          : 1/0/0
   NCQ status                               : Enabled
   Queue Depth                              : Automatic
   Monitor and Performance Delay            : 60 minutes
   Elevator Sort                            : Enabled
   Degraded Mode Performance Optimization   : Disabled
   Latency                                  : Default
   Statistics data collection mode          : Disabled
   Global Physical Device Write Cache Policy: Disabled
   maxCache RAID5 WriteBack Enabled         : Disabled
   maxCache Version                         : Not Applicable
   Post Prompt Timeout                      : 15 seconds
   Boot Controller                          : False
   Primary Boot Volume                      : None
   Secondary Boot Volume                    : None
   Driver Name                              : aacraid
   Driver Supports SSD Smart Path           : No
   Manufacturing Part Number                : Not Applicable
   Manufacturing Spare Part Number          : Not Applicable
   Manufacturing Wellness Log               : Not Applicable
   --------------------------------------------------------
   RAID Properties
   --------------------------------------------------------
   Automatic Failover                       : Disabled
   Background consistency check             : Idle
   Consistency Check Delay                  : 3 seconds
   Parallel Consistency Check Supported     : Enabled
   Parallel Consistency Check Count         : 1
   Inconsistency Repair Policy              : Disabled
   Consistency Check Inconsistency Notify   : Disabled
   Rebuild Priority                         : High
   Expand Priority                          : Medium
   --------------------------------------------------------
   Controller Version Information
   --------------------------------------------------------
   Firmware                                 : 3.52[0]
   Driver                                   : 1.2.1.50877
   Hardware Revision                        : B
   --------------------------------------------------------
   Temperature Sensors Information
   --------------------------------------------------------
   Sensor ID                                : 0
   Current Value                            : 32 deg C
   Max Value Since Powered On               : 33 deg C
   Location                                 : Inlet Ambient

   Sensor ID                                : 1
   Current Value                            : 48 deg C
   Max Value Since Powered On               : 49 deg C
   Location                                 : ASIC


----------------------------------------------------------------------
Array Information
----------------------------------------------------------------------
Array Number 0
   Name                                     : A
   Status                                   : Ok
   Interface                                : SAS
   Total Size                               : 571904 MB
   Unused Size                              : 0 MB
   Block Size                               : 512 Bytes
   Array Utilization                        : 100.00% Used, 0.00% Unused
   Type                                     : Data
   Transformation Status                    : Not Applicable
   Spare Rebuild Mode                       : Dedicated
   --------------------------------------------------------
   Array Logical Device Information
   --------------------------------------------------------
   Logical 0                                : Optimal (1, Data, 286136 MB) LogicalDrv 0
   --------------------------------------------------------
   Array Physical Device Information
   --------------------------------------------------------
   Device 26                                : Present (286168MB, SAS, HDD, Connector:0, Enclosure:1, Slot:28)         0AGRGZTH
   Device 27                                : Present (286168MB, SAS, HDD, Connector:0, Enclosure:1, Slot:29)         0AGRJ8ZH


----------------------------------------------------------------------
Logical device information
----------------------------------------------------------------------
Logical Device number 0
   Logical Device name                      : LogicalDrv 0
   Disk Name                                : /dev/sda
   Array                                    : 0
   RAID level                               : 1
   Status of Logical Device                 : Optimal
   Size                                     : 286136 MB
   Full Stripe Size                         : 256 KB
   Interface Type                           : Serial Attached SCSI
   Device Type                              : Data
   Boot Type                                : None
   Heads                                    : 255
   Sectors Per Track                        : 32
   Cylinders                                : 65535
   Mount Points                             : Not Applicable
   LD Acceleration Method                   : None
   Volume Unique Identifier                 : 600508B1001C2EB4A2D267CB8008C226
   --------------------------------------------------------
   Array Physical Device Information
   --------------------------------------------------------
   Device 26                                : Present (286168MB, SAS, HDD, Connector:0, Enclosure:1, Slot:28)         0AGRGZTH
   Device 27                                : Present (286168MB, SAS, HDD, Connector:0, Enclosure:1, Slot:29)         0AGRJ8ZH


----------------------------------------------------------------------
Physical Device information
----------------------------------------------------------------------
   Channel #0:
      Device #8
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdb
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,8(8:0)
         Reported Location                  : Enclosure 1, Slot 0(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8G72GUD
         World-wide name                    : 53CE5A63C6BE9000
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 43 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : BF1324C6FF7FAF02FF7FA3DAFFFF4565
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #9
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdc
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,9(9:0)
         Reported Location                  : Enclosure 1, Slot 1(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8G9N2BN
         World-wide name                    : 53CE5A63C6BE9001
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 44 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3F3477B83FAD5F613FAD58513F21DB55
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #10
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdd
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,10(10:0)
         Reported Location                  : Enclosure 1, Slot 2(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8GADKYN
         World-wide name                    : 53CE5A63C6BE9002
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 43 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3FCCE374FF7B3FD1FF7B3E99FFBBE0D9
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #11
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sde
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,11(11:0)
         Reported Location                  : Enclosure 1, Slot 3(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8GA18MD
         World-wide name                    : 53CE5A63C6BE9003
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 30 deg C
         Maximum Temperature                : 45 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3F7AA1FA7FFE23DB3FFE22AB3F66A58C
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #12
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdf
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,12(12:0)
         Reported Location                  : Enclosure 1, Slot 4(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : ST6000NM0115-1YZ
         Firmware                           : SN05    
         Serial number                      : ZADAQJMV
         World-wide name                    : 53CE5A63C6BE9004
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 27 deg C
         Maximum Temperature                : 44 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : True
         Drive Unique ID                    : 3F5C016CBFBFFF0EBFBC578EBF3C9A92
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #13
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdg
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,13(13:0)
         Reported Location                  : Enclosure 1, Slot 5(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8GABGVN
         World-wide name                    : 53CE5A63C6BE9005
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 44 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : FF817C9FFFFF7F22FFEF67F2FF6D0AED
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #14
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdh
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,14(14:0)
         Reported Location                  : Enclosure 1, Slot 6(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8G9WLPD
         World-wide name                    : 53CE5A63C6BE9006
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 43 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3FD465E37FFB4BB57F7B4A453FC2CC26
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #15
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdi
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,15(15:0)
         Reported Location                  : Enclosure 1, Slot 7(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8GA39KD
         World-wide name                    : 53CE5A63C6BE9007
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 43 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3FAC8A6CFFBFBD47FF1BB5FF7F9B9772
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #16
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdj
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,16(16:0)
         Reported Location                  : Enclosure 1, Slot 8(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8G9WBKD
         World-wide name                    : 53CE5A63C6BE9008
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 43 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3F0C7AABFFFF3D9DFF7E3CB57FFABEA6
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #17
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdk
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,17(17:0)
         Reported Location                  : Enclosure 1, Slot 9(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8G6A3JD
         World-wide name                    : 53CE5A63C6BE9009
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 44 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3F0F8B66FF7FA714FF7BA374BFFB6540
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #18
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdl
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,18(18:0)
         Reported Location                  : Enclosure 1, Slot 10(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8GAA8WD
         World-wide name                    : 53CE5A63C6BE900A
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 44 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3FB2F3BEFFDE9B377F5E9B0F7F9E7DEB
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #19
         Device is a Hard drive
         State                              : Ready
         Drive has stale RIS data           : False
         Disk Name                          : /dev/sdm
         Block Size                         : 512 Bytes
         Physical Block Size                : 4K Bytes
         Transfer Speed                     : SATA 6.0 Gb/s
         Reported Channel,Device(T:L)       : 0,19(19:0)
         Reported Location                  : Enclosure 1, Slot 11(Connector 0:CN0)
         Vendor                             : ATA     
         Model                              : HGST HUS726060AL
         Firmware                           : APBDTAF0
         Serial number                      : K8G89X1D
         World-wide name                    : 53CE5A63C6BE900B
         Reserved Size                      : 32768 KB
         Used Size                          : 0 MB
         Unused Size                        : 5723134 MB
         Total Size                         : 5723166 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         NCQ supported                      : Supported
         NCQ status                         : Enabled
         Boot Type                          : None
         Rotational Speed                   : 7200 RPM
         Current Temperature                : 29 deg C
         Maximum Temperature                : 43 deg C
         Threshold Temperature              : 65 deg C
         PHY Count                          : 1
         Drive Configuration Type           : Unassigned
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : True
         Sanitize Erase Support             : False
         Drive Unique ID                    : 3F7C24CBFFFF39F3FFFF392B7F6B1C9D
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : 6 Gbps 
            Maximum Link Rate               : 6 Gbps 

      ----------------------------------------------------------------
      Device Error Counters                 
      ----------------------------------------------------------------
         Aborted Commands                   : 0
         Bad Target Errors                  : 0
         Ecc Recovered Read Errors          : 0
         Failed Read Recovers               : 0
         Failed Write Recovers              : 0
         Format Errors                      : 0
         Hardware Errors                    : 0
         Hard Read Errors                   : 0
         Hard Write Errors                  : 0
         Hot Plug Count                     : 0
         Media Failures                     : 0
         Not Ready Errors                   : 0
         Other Time Out Errors              : 0
         Predictive Failures                : 0
         Retry Recovered Read Errors        : 0
         Retry Recovered Write Errors       : 0
         Scsi Bus Faults                    : 0

      Device #26
         Device is a Hard drive
         State                              : Online
         Drive has stale RIS data           : False
         Block Size                         : 512 Bytes
         Physical Block Size                : 512 Bytes
         Transfer Speed                     : SAS 12.0 Gb/s
         Reported Channel,Device(T:L)       : 0,26(26:0)
         Reported Location                  : Enclosure 1, Slot 28(Connector 0:CN0)
         Array                              : 0
         Vendor                             : HGST    
         Model                              : HUC101830CSS200 
         Firmware                           : ADB0
         Serial number                      : 0AGRGZTH
         World-wide name                    : 5000CCA07C28E1C9
         Reserved Size                      : 32768 KB
         Used Size                          : 286136 MB
         Unused Size                        : 0 MB
         Total Size                         : 286168 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         Boot Type                          : None
         Rotational Speed                   : 10500 RPM
         Current Temperature                : 34 deg C
         Maximum Temperature                : Not Applicable
         Threshold Temperature              : 85 deg C
         PHY Count                          : 2
         Drive Configuration Type           : Data
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : False
         Sanitize Erase Support             : False
         Drive Unique ID                    : 5000CCA07C28E1C8619300085000CCA0
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : Reserved [11] 
            Maximum Link Rate               : Reserved [11] 
         Phy #1
            Negotiated Link Rate            : unknown 
            Maximum Link Rate               : Reserved [11] 


      Device #27
         Device is a Hard drive
         State                              : Online
         Drive has stale RIS data           : False
         Block Size                         : 512 Bytes
         Physical Block Size                : 512 Bytes
         Transfer Speed                     : SAS 12.0 Gb/s
         Reported Channel,Device(T:L)       : 0,27(27:0)
         Reported Location                  : Enclosure 1, Slot 29(Connector 0:CN0)
         Array                              : 0
         Vendor                             : HGST    
         Model                              : HUC101830CSS200 
         Firmware                           : ADB0
         Serial number                      : 0AGRJ8ZH
         World-wide name                    : 5000CCA07C28F541
         Reserved Size                      : 32768 KB
         Used Size                          : 286136 MB
         Unused Size                        : 0 MB
         Total Size                         : 286168 MB
         S.M.A.R.T.                         : No
         S.M.A.R.T. warnings                : 0
         SSD                                : No
         Boot Type                          : None
         Rotational Speed                   : 10500 RPM
         Current Temperature                : 35 deg C
         Maximum Temperature                : Not Applicable
         Threshold Temperature              : 85 deg C
         PHY Count                          : 2
         Drive Configuration Type           : Data
         Mount Point(s)                     : Not Mounted
         Drive Exposed to OS                : False
         Sanitize Erase Support             : False
         Drive Unique ID                    : 5000CCA07C28F540619300085000CCA0
         Last Failure Reason                : No Failure
      ----------------------------------------------------------------
      Device Phy Information                
      ----------------------------------------------------------------
         Phy #0
            Negotiated Link Rate            : Reserved [11] 
            Maximum Link Rate               : Reserved [11] 
         Phy #1
            Negotiated Link Rate            : unknown 
            Maximum Link Rate               : Reserved [11] 


   Channel #2:
      Device #0
         Device is an Enclosure Services Device
         Reported Channel,Device(T:L)       : 2,0(0:0)
         Reported Location                  : Connector 0, Enclosure 1
         Enclosure ID                       : 1
         Enclosure Logical Identifier       : 53CE5A63C6BE907E
         Expander ID                        : 0
         Expander SAS Address               : 53CE5A63C6BE907F
         Type                               : SES2
         Vendor                             : H3C-Exp
         Model                              : SXP 36x12G
         Firmware                           : RevB
         Status of Enclosure Services Device
            Speaker status                  : Not Available
   Channel #2:
      Device #3
         Device is an Enclosure Services Device
         Reported Channel,Device(T:L)       : 2,3(3:0)
         Enclosure ID                       : 0
         Enclosure Logical Identifier       : 50123456789ABC00
         Type                               : SES2
         Vendor                             : ADAPTEC 
         Model                              : Virtual SGPIO
         Firmware                           : 3.52
         Status of Enclosure Services Device
            Speaker status                  : Not Available


----------------------------------------------------------------------
Connector information
----------------------------------------------------------------------
Connector #0
   Connector Name                           : CN0
   Functional Mode                          : Mixed
   Connector Location                       : Internal
   SAS Address                              : 50123456789ABC00


Connector #1
   Connector Name                           : CN1
   Functional Mode                          : Mixed
   Connector Location                       : Internal
   SAS Address                              : 50123456789ABC04

