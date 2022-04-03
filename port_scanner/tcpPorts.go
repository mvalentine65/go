type TCP int32:

const (
    FTP-DATA TCP= 20
    FTP         = 21
    SSH         = 22
    TELNET      = 23
    SMTP        = 25
    NAMESERVER  = 42
    DNS         = 53
    HTTP        = 80
    KERBEROS    = 88
    POP3        = 110
    SFTP        = 115
    NNTP        = 119
    RPC         = 135
    NETBIOS     = 137
    IMAP4       = 143
    SNMP        = 161
    SNMPTRAP    = 162
    BGP         = 179
    LDAP        = 389
    HTTPS       = 443
    SMB         = 445
    PRINTER     = 515
    IPP         = 631
    FTPS-DATA   = 989
    FTPS        = 990
)

func (port TCP) String() string {
    switch port {
    case FTP-DATA:
        return "ftp-data"
    case FTP:
        return "ftp"
    case SHH:
        return "ssh"
    case TELNET:
        return "telnet"
    case SMTP:
        return "smtp"
    case NAMESERVER:
        return "nameserver"
    case DNS:
        return "dns"
    case HTTP:
        return "http"
    case KERBEROS:
        return "kerberos"
    case POP3:
        return "pop3"
    case SFTP:
        return "sftp"
    case NNTP:
        return "nntp"
    case RPC:
        return "rpc"
    case NETBIOS:
        return "netbios"
    case IMAP4:
        return "imap4"
    case SNMP:
        return "snmp"
    case SNMPTRAP:
        return "snmp-trap"
    case BGP:
        return "bgp"
    case LDAP:
        return "ldap"
    case HTTPS:
        return "https"
    case SMB:
        return "smb"
    case PRINTER:
        return "printer"
    case IPP:
        return "ipp"
    case FTPS-DATA:
        return "ftps-data"
    case FTPS:
        return "ftps"
    }
    return "unknown"
}

