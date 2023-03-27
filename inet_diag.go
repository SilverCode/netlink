package netlink

// INET_DIAG constatns
const (
	INET_DIAG_NONE = iota
	INET_DIAG_MEMINFO
	INET_DIAG_INFO
	INET_DIAG_VEGASINFO
	INET_DIAG_CONG
	INET_DIAG_TOS
	INET_DIAG_TCLASS
	INET_DIAG_SKMEMINFO
	INET_DIAG_SHUTDOWN
	INET_DIAG_DCTCPINFO
	INET_DIAG_PROTOCOL
	INET_DIAG_SKV6ONLY
	INET_DIAG_LOCALS
	INET_DIAG_PEERS
	INET_DIAG_PAD
	INET_DIAG_MARK
	INET_DIAG_BBRINFO
	INET_DIAG_CLASS_ID
	INET_DIAG_MD5SIG
	INET_DIAG_ULP_INFO
	INET_DIAG_SK_BPF_STORAGES
	INET_DIAG_CGROUP_ID
	INET_DIAG_SOCKOPT
	INET_DIAG_MAX
)

type InetDiagTCPInfoResp struct {
	InetDiagMsg *Socket
	TCPInfo     *TCPInfo
	TCPBBRInfo  *TCPBBRInfo
	CGroupID    uint64
}
