# Импорт модуля ServerManager
Import-Module ServerManager
# Установка роли AD DS со всеми зависимыми компонентами
Add-WindowsFeature –Name AD-Domain-Services –IncludeAllSubFeature –IncludeManagementTools
Import-Module ADDSDeployment
Install-ADDSForest `
-CreateDnsDelegation:$false `
-DatabasePath "C:\Windows\NTDS" `
-DomainMode "WinThreshold" `
-DomainName "test.local" `
-DomainNetbiosName "TEST" `
-ForestMode "WinThreshold" `
-InstallDns:$true `
-LogPath "C:\Windows\NTDS" `
-NoRebootOnCompletion:$false `
-SysvolPath "C:\Windows\SYSVOL" `
-Force:$true
 Add-DnsServerPrimaryZone -NetworkId '192.168.5.0/27'
 -ReplicationScope Domain
 Add-DnsServerResourceRecordPtr -Name'2'
 -ZoneName '5.168.192.in-addr.arpa'
 -AgeRecord
 -PtrDomainName 'FreshTest.local'
# Import-Module ServerManager
# Add-WindowsFeature -name DHCP -IncludeManagementTools
# Add-DHCPServerV4Scope -Name Pool1 -StartRange 192.168.0.5
# -EndRange 192.168.0.100 -SubnetMask 255.255.255.0
# -State Active
# Set-DHCPServerV4OptionValue -DNSServer 192.168.0.1
# -DnsDomain test.local -Router 192.168.0.1
DSadd group 'CN=group1,Cn=Users,DC=test,DC=local'
-samid group1 -secgrp yes -scope g
DSadd group 'CN=group2,Cn=Users,DC=test,DC=local'
-samid group2 -secgrp yes -scope g
DSadd group 'CN=group3,Cn=Users,DC=test,DC=local'
-samid group3 -secgrp yes -scope g
DSadd user CN=Иванов,Cn=Users,DC=test,DC=local
-samid Иванов -pwd 93Noob93
DSadd user CN=Петров,Cn=Users,DC=test,DC=local
-samid Петров -pwd 93Noob93
DSadd user CN=Сидоров,Cn=Users,DC=test,DC=local
-samid Сидоров -pwd 93Noob93
DSadd user CN=Козлов,Cn=Users,DC=test,DC=local
-samid Козлов -pwd 93Noob93
DSadd user CN=Рембо,Cn=Users,DC=test,DC=local
-samid Рембо -pwd 93Noob93
Add-ADGroupMember 'group1' group1
Add-ADGroupMember 'group1' group1
Add-ADGroupMember 'group1' group1
Add-ADGroupMember 'group2' group2
Add-ADGroupMember 'group2' group2
Add-ADGroupMember 'group3' group3
Add-ADGroupMember 'group3' group3
