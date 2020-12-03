package openbmc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/bmc-toolbox/bmclib/cfgresources"
	"github.com/bmc-toolbox/bmclib/devices"

	"github.com/go-logr/logr"
)

type OpenBmc struct {
	ip         string
	username   string
	password   string
	httpClient *http.Client
	ctx        context.Context
	log        logr.Logger
}

func (b *OpenBmc) Bios(cfg *cfgresources.Bios) (err error) {
	return err
}

func (b *OpenBmc) BiosVersion() (version string, err error) {
	return version, err
}

func (b *OpenBmc) CPU() (cpu string, cpuCount int, coreCount int, hyperthreadCount int, err error) {
	return cpu, cpuCount, coreCount, hyperthreadCount, err
}

func (b *OpenBmc) ChassisSerial() (serial string, err error) {
	return serial, err
}

func (b *OpenBmc) CheckCredentials() (err error) {
	return err
}

func (b *OpenBmc) Close() (err error) {
	return err
}

func (b *OpenBmc) CurrentHTTPSCert() ([]*x509.Certificate, bool, error) {

	dialer := &net.Dialer{
		Timeout: time.Duration(10) * time.Second,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", b.ip+":"+"443", &tls.Config{InsecureSkipVerify: true})

	if err != nil {
		return []*x509.Certificate{{}}, false, err
	}

	defer conn.Close()

	return conn.ConnectionState().PeerCertificates, false, nil

}

func (b *OpenBmc) Disks() (disks []*devices.Disk, err error) {
	return disks, err
}

func (b *OpenBmc) GenerateCSR(cert *cfgresources.HTTPSCertAttributes) ([]byte, error) {
	return []byte{}, nil
}

func (b *OpenBmc) HardwareType() (model string) {
	return model
}

func (b *OpenBmc) IsBlade() (isBlade bool, err error) {
	return isBlade, err
}

func (b *OpenBmc) IsOn() (status bool, err error) {
	return status, err
}

func (b *OpenBmc) Ldap(cfgLdap *cfgresources.Ldap) error {
	return nil
}

func (b *OpenBmc) LdapGroup(cfgGroup []*cfgresources.LdapGroup, cfgLdap *cfgresources.Ldap) (err error) {
	return err
}

func (b *OpenBmc) License() (name string, licType string, err error) {
	return name, licType, err
}

func (b *OpenBmc) Memory() (mem int, err error) {
	return mem, err
}

func (b *OpenBmc) Model() (model string, err error) {
	return model, err
}

func (b *OpenBmc) Name() (name string, err error) {
	return name, err
}

func (b *OpenBmc) Network(cfg *cfgresources.Network) (reset bool, err error) {
	return reset, err
}

func (b *OpenBmc) Nics() (nics []*devices.Nic, err error) {
	return nics, err
}

func (b *OpenBmc) Ntp(cfg *cfgresources.Ntp) (err error) {
	return err
}

func (b *OpenBmc) Power(cfg *cfgresources.Power) (err error) {
	return err
}

func (b *OpenBmc) PowerCycle() (status bool, err error) {
	return status, err
}

func (b *OpenBmc) PowerCycleBmc() (status bool, err error) {
	return status, err
}

func (b *OpenBmc) PowerKw() (power float64, err error) {
	return power, err
}

func (b *OpenBmc) PowerOn() (status bool, err error) {
	return status, err
}

func (b *OpenBmc) PowerOff() (status bool, err error) {
	return status, err
}

func (b *OpenBmc) PowerState() (state string, err error) {
	return state, err
}

func (b *OpenBmc) PxeOnce() (status bool, err error) {
	return status, err
}

func (b *OpenBmc) Resources() []string {
	return []string{}
}

func (b *OpenBmc) Screenshot() (response []byte, extension string, err error) {
	return response, extension, err
}

func (b *OpenBmc) Serial() (serial string, err error) {
	return serial, err
}

func (b *OpenBmc) ServerSnapshot() (server interface{}, err error) {
	return server, err
}

func (b *OpenBmc) SetLicense(cfg *cfgresources.License) (err error) {
	return err
}

func (b *OpenBmc) Slot() (slot int, err error) {
	return slot, err
}

func (b *OpenBmc) Status() (health string, err error) {
	return health, err
}

func (b *OpenBmc) Syslog(cfg *cfgresources.Syslog) (err error) {
	return err
}

func (b *OpenBmc) TempC() (temp int, err error) {
	return temp, err
}

func (b *OpenBmc) UpdateCredentials(username string, password string) {
	b.username = username
	b.password = password
}

func (b *OpenBmc) UpdateFirmware(source, file string) (status bool, err error) {
	return true, fmt.Errorf("NYI")
}

func (b *OpenBmc) UploadHTTPSCert(cert []byte, certFileName string, key []byte, keyFileName string) (bool, error) {
	return false, fmt.Errorf("NYI")
}

func (b *OpenBmc) User(users []*cfgresources.User) (err error) {
	return err
}

func (b *OpenBmc) Vendor() (vendor string) {
	return "OpenBMC"
}

func (b *OpenBmc) Version() (bmcVersion string, err error) {
	return bmcVersion, err
}

func New(ctx context.Context, ip string, username string, password string, log logr.Logger) (obmc *OpenBmc, err error) {
	var _ devices.Bmc = &OpenBmc{}
	return &OpenBmc{
		ip:       ip,
		username: username,
		password: password,
		ctx:      ctx,
		log:      log,
	}, err
}
