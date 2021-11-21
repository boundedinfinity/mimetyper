//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package file_extention

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type FileExtention string
type FileExtentions []FileExtention

func Slice(es ...FileExtention) FileExtentions {
	var s FileExtentions

	for _, e := range es {
		s = append(s, e)
	}

	return s
}

const (
	_123        FileExtention = ".123"
	_3dml       FileExtention = ".3dml"
	_3g2        FileExtention = ".3g2"
	_3gp        FileExtention = ".3gp"
	_7z         FileExtention = ".7z"
	Aab         FileExtention = ".aab"
	Aac         FileExtention = ".aac"
	Aam         FileExtention = ".aam"
	Aas         FileExtention = ".aas"
	Abw         FileExtention = ".abw"
	Ac          FileExtention = ".ac"
	Acc         FileExtention = ".acc"
	Ace         FileExtention = ".ace"
	Acu         FileExtention = ".acu"
	Adp         FileExtention = ".adp"
	Aep         FileExtention = ".aep"
	Afp         FileExtention = ".afp"
	Ahead       FileExtention = ".ahead"
	Ai          FileExtention = ".ai"
	Aif         FileExtention = ".aif"
	Air         FileExtention = ".air"
	Ait         FileExtention = ".ait"
	Ami         FileExtention = ".ami"
	Apk         FileExtention = ".apk"
	Application FileExtention = ".application"
	Apr         FileExtention = ".apr"
	Arc         FileExtention = ".arc"
	Asf         FileExtention = ".asf"
	Aso         FileExtention = ".aso"
	Atc         FileExtention = ".atc"
	Atom        FileExtention = ".atom"
	Atomcat     FileExtention = ".atomcat"
	Atomsvc     FileExtention = ".atomsvc"
	Atx         FileExtention = ".atx"
	Au          FileExtention = ".au"
	Avi         FileExtention = ".avi"
	Aw          FileExtention = ".aw"
	Azf         FileExtention = ".azf"
	Azs         FileExtention = ".azs"
	Azw         FileExtention = ".azw"
	Bcpio       FileExtention = ".bcpio"
	Bdf         FileExtention = ".bdf"
	Bdm         FileExtention = ".bdm"
	Bed         FileExtention = ".bed"
	Bh2         FileExtention = ".bh2"
	Bin         FileExtention = ".bin"
	Bmi         FileExtention = ".bmi"
	Bmp         FileExtention = ".bmp"
	Box         FileExtention = ".box"
	Btif        FileExtention = ".btif"
	Bz          FileExtention = ".bz"
	Bz2         FileExtention = ".bz2"
	C           FileExtention = ".c"
	C11amc      FileExtention = ".c11amc"
	C11amz      FileExtention = ".c11amz"
	C4g         FileExtention = ".c4g"
	Cab         FileExtention = ".cab"
	Car         FileExtention = ".car"
	Cat         FileExtention = ".cat"
	Ccxml       FileExtention = ".ccxml"
	Cda         FileExtention = ".cda"
	Cdbcmsg     FileExtention = ".cdbcmsg"
	Cdkey       FileExtention = ".cdkey"
	Cdmia       FileExtention = ".cdmia"
	Cdmic       FileExtention = ".cdmic"
	Cdmid       FileExtention = ".cdmid"
	Cdmio       FileExtention = ".cdmio"
	Cdmiq       FileExtention = ".cdmiq"
	Cdx         FileExtention = ".cdx"
	Cdxml       FileExtention = ".cdxml"
	Cdy         FileExtention = ".cdy"
	Cer         FileExtention = ".cer"
	Cgm         FileExtention = ".cgm"
	Chat        FileExtention = ".chat"
	Chm         FileExtention = ".chm"
	Chrt        FileExtention = ".chrt"
	Cif         FileExtention = ".cif"
	Cii         FileExtention = ".cii"
	Cil         FileExtention = ".cil"
	Cla         FileExtention = ".cla"
	Class       FileExtention = ".class"
	Clkk        FileExtention = ".clkk"
	Clkp        FileExtention = ".clkp"
	Clkt        FileExtention = ".clkt"
	Clkw        FileExtention = ".clkw"
	Clkx        FileExtention = ".clkx"
	Clp         FileExtention = ".clp"
	Cmc         FileExtention = ".cmc"
	Cmdf        FileExtention = ".cmdf"
	Cml         FileExtention = ".cml"
	Cmp         FileExtention = ".cmp"
	Cmx         FileExtention = ".cmx"
	Cod         FileExtention = ".cod"
	Cpio        FileExtention = ".cpio"
	Cpt         FileExtention = ".cpt"
	Crd         FileExtention = ".crd"
	Crl         FileExtention = ".crl"
	Cryptonote  FileExtention = ".cryptonote"
	Csh         FileExtention = ".csh"
	Csml        FileExtention = ".csml"
	Csp         FileExtention = ".csp"
	Css         FileExtention = ".css"
	Csv         FileExtention = ".csv"
	Cu          FileExtention = ".cu"
	Curl        FileExtention = ".curl"
	Cww         FileExtention = ".cww"
	Dae         FileExtention = ".dae"
	Daf         FileExtention = ".daf"
	Davmount    FileExtention = ".davmount"
	Dcurl       FileExtention = ".dcurl"
	Dd2         FileExtention = ".dd2"
	Ddd         FileExtention = ".ddd"
	Deb         FileExtention = ".deb"
	Der         FileExtention = ".der"
	Dfac        FileExtention = ".dfac"
	Dir         FileExtention = ".dir"
	Dis         FileExtention = ".dis"
	Djvu        FileExtention = ".djvu"
	Dmg         FileExtention = ".dmg"
	Dna         FileExtention = ".dna"
	Doc         FileExtention = ".doc"
	Docm        FileExtention = ".docm"
	Docx        FileExtention = ".docx"
	Dotm        FileExtention = ".dotm"
	Dotx        FileExtention = ".dotx"
	Dp          FileExtention = ".dp"
	Dpg         FileExtention = ".dpg"
	Dra         FileExtention = ".dra"
	Dsc         FileExtention = ".dsc"
	Dssc        FileExtention = ".dssc"
	Dtb         FileExtention = ".dtb"
	Dtd         FileExtention = ".dtd"
	Dts         FileExtention = ".dts"
	Dtshd       FileExtention = ".dtshd"
	Dvi         FileExtention = ".dvi"
	Dwf         FileExtention = ".dwf"
	Dwg         FileExtention = ".dwg"
	Dxf         FileExtention = ".dxf"
	Dxp         FileExtention = ".dxp"
	Ecelp4800   FileExtention = ".ecelp4800"
	Ecelp7470   FileExtention = ".ecelp7470"
	Ecelp9600   FileExtention = ".ecelp9600"
	Edm         FileExtention = ".edm"
	Edx         FileExtention = ".edx"
	Efif        FileExtention = ".efif"
	Ei6         FileExtention = ".ei6"
	Eml         FileExtention = ".eml"
	Emma        FileExtention = ".emma"
	Eol         FileExtention = ".eol"
	Eot         FileExtention = ".eot"
	Epub        FileExtention = ".epub"
	Es          FileExtention = ".es"
	Es3         FileExtention = ".es3"
	Esf         FileExtention = ".esf"
	Etx         FileExtention = ".etx"
	Exe         FileExtention = ".exe"
	Exi         FileExtention = ".exi"
	Ext         FileExtention = ".ext"
	Ez2         FileExtention = ".ez2"
	Ez3         FileExtention = ".ez3"
	F           FileExtention = ".f"
	F4v         FileExtention = ".f4v"
	Fbs         FileExtention = ".fbs"
	Fcs         FileExtention = ".fcs"
	Fdf         FileExtention = ".fdf"
	FeLaunch    FileExtention = ".fe_launch"
	Fg5         FileExtention = ".fg5"
	Fh          FileExtention = ".fh"
	Fig         FileExtention = ".fig"
	Fli         FileExtention = ".fli"
	Flo         FileExtention = ".flo"
	Flv         FileExtention = ".flv"
	Flw         FileExtention = ".flw"
	Flx         FileExtention = ".flx"
	Fly         FileExtention = ".fly"
	Fm          FileExtention = ".fm"
	Fnc         FileExtention = ".fnc"
	Fpx         FileExtention = ".fpx"
	Fsc         FileExtention = ".fsc"
	Fst         FileExtention = ".fst"
	Ftc         FileExtention = ".ftc"
	Fti         FileExtention = ".fti"
	Fvt         FileExtention = ".fvt"
	Fxp         FileExtention = ".fxp"
	Fzs         FileExtention = ".fzs"
	G2w         FileExtention = ".g2w"
	G3          FileExtention = ".g3"
	G3w         FileExtention = ".g3w"
	Gac         FileExtention = ".gac"
	Gdl         FileExtention = ".gdl"
	Geo         FileExtention = ".geo"
	Gex         FileExtention = ".gex"
	Ggb         FileExtention = ".ggb"
	Ggt         FileExtention = ".ggt"
	Ghf         FileExtention = ".ghf"
	Gif         FileExtention = ".gif"
	Gim         FileExtention = ".gim"
	Gmx         FileExtention = ".gmx"
	Gnumeric    FileExtention = ".gnumeric"
	Gph         FileExtention = ".gph"
	Gpx         FileExtention = ".gpx"
	Gqf         FileExtention = ".gqf"
	Gram        FileExtention = ".gram"
	Grv         FileExtention = ".grv"
	Grxml       FileExtention = ".grxml"
	Gsf         FileExtention = ".gsf"
	Gtar        FileExtention = ".gtar"
	Gtm         FileExtention = ".gtm"
	Gtw         FileExtention = ".gtw"
	Gv          FileExtention = ".gv"
	Gxt         FileExtention = ".gxt"
	Gz          FileExtention = ".gz"
	H261        FileExtention = ".h261"
	H263        FileExtention = ".h263"
	H264        FileExtention = ".h264"
	Hal         FileExtention = ".hal"
	Hbci        FileExtention = ".hbci"
	Hdf         FileExtention = ".hdf"
	Hlp         FileExtention = ".hlp"
	Hpgl        FileExtention = ".hpgl"
	Hpid        FileExtention = ".hpid"
	Hps         FileExtention = ".hps"
	Hqx         FileExtention = ".hqx"
	Htke        FileExtention = ".htke"
	Html        FileExtention = ".html"
	Hvd         FileExtention = ".hvd"
	Hvp         FileExtention = ".hvp"
	Hvs         FileExtention = ".hvs"
	I2g         FileExtention = ".i2g"
	Icc         FileExtention = ".icc"
	Ice         FileExtention = ".ice"
	Ico         FileExtention = ".ico"
	Ics         FileExtention = ".ics"
	Ief         FileExtention = ".ief"
	Ifm         FileExtention = ".ifm"
	Igl         FileExtention = ".igl"
	Igm         FileExtention = ".igm"
	Igs         FileExtention = ".igs"
	Igx         FileExtention = ".igx"
	Iif         FileExtention = ".iif"
	Imp         FileExtention = ".imp"
	Ims         FileExtention = ".ims"
	Ipfix       FileExtention = ".ipfix"
	Ipk         FileExtention = ".ipk"
	Irm         FileExtention = ".irm"
	Irp         FileExtention = ".irp"
	Itp         FileExtention = ".itp"
	Ivp         FileExtention = ".ivp"
	Ivu         FileExtention = ".ivu"
	Jad         FileExtention = ".jad"
	Jam         FileExtention = ".jam"
	Jar         FileExtention = ".jar"
	Java        FileExtention = ".java"
	Jisp        FileExtention = ".jisp"
	Jlt         FileExtention = ".jlt"
	Jnlp        FileExtention = ".jnlp"
	Joda        FileExtention = ".joda"
	Jpeg        FileExtention = ".jpeg"
	Jpg         FileExtention = ".jpg"
	Jpgv        FileExtention = ".jpgv"
	Jpm         FileExtention = ".jpm"
	Js          FileExtention = ".js"
	Json        FileExtention = ".json"
	Jsonld      FileExtention = ".jsonld"
	Karbon      FileExtention = ".karbon"
	Kfo         FileExtention = ".kfo"
	Kia         FileExtention = ".kia"
	Kml         FileExtention = ".kml"
	Kmz         FileExtention = ".kmz"
	Kne         FileExtention = ".kne"
	Kon         FileExtention = ".kon"
	Kpr         FileExtention = ".kpr"
	Ksp         FileExtention = ".ksp"
	Ktx         FileExtention = ".ktx"
	Ktz         FileExtention = ".ktz"
	Kwd         FileExtention = ".kwd"
	Lasxml      FileExtention = ".lasxml"
	Latex       FileExtention = ".latex"
	Lbd         FileExtention = ".lbd"
	Lbe         FileExtention = ".lbe"
	Les         FileExtention = ".les"
	Link66      FileExtention = ".link66"
	Lrm         FileExtention = ".lrm"
	Ltf         FileExtention = ".ltf"
	Lvp         FileExtention = ".lvp"
	Lwp         FileExtention = ".lwp"
	M21         FileExtention = ".m21"
	M3u         FileExtention = ".m3u"
	M3u8        FileExtention = ".m3u8"
	M4v         FileExtention = ".m4v"
	Ma          FileExtention = ".ma"
	Mads        FileExtention = ".mads"
	Mag         FileExtention = ".mag"
	Mathml      FileExtention = ".mathml"
	Mbk         FileExtention = ".mbk"
	Mbox        FileExtention = ".mbox"
	Mc1         FileExtention = ".mc1"
	Mcd         FileExtention = ".mcd"
	Mcurl       FileExtention = ".mcurl"
	Mdb         FileExtention = ".mdb"
	Mdi         FileExtention = ".mdi"
	Meta4       FileExtention = ".meta4"
	Mets        FileExtention = ".mets"
	Mfm         FileExtention = ".mfm"
	Mgp         FileExtention = ".mgp"
	Mgz         FileExtention = ".mgz"
	Mid         FileExtention = ".mid"
	Midi        FileExtention = ".midi"
	Mif         FileExtention = ".mif"
	Mj2         FileExtention = ".mj2"
	Mlp         FileExtention = ".mlp"
	Mmd         FileExtention = ".mmd"
	Mmf         FileExtention = ".mmf"
	Mmr         FileExtention = ".mmr"
	Mny         FileExtention = ".mny"
	Mods        FileExtention = ".mods"
	Movie       FileExtention = ".movie"
	Mp4         FileExtention = ".mp4"
	Mp4a        FileExtention = ".mp4a"
	Mpc         FileExtention = ".mpc"
	Mpeg        FileExtention = ".mpeg"
	Mpga        FileExtention = ".mpga"
	Mpkg        FileExtention = ".mpkg"
	Mpm         FileExtention = ".mpm"
	Mpn         FileExtention = ".mpn"
	Mpp         FileExtention = ".mpp"
	Mpy         FileExtention = ".mpy"
	Mqy         FileExtention = ".mqy"
	Mrc         FileExtention = ".mrc"
	Mrcx        FileExtention = ".mrcx"
	Mscml       FileExtention = ".mscml"
	Mseq        FileExtention = ".mseq"
	Msf         FileExtention = ".msf"
	Msh         FileExtention = ".msh"
	Msl         FileExtention = ".msl"
	Msty        FileExtention = ".msty"
	Mts         FileExtention = ".mts"
	Mus         FileExtention = ".mus"
	Musicxml    FileExtention = ".musicxml"
	Mvb         FileExtention = ".mvb"
	Mwf         FileExtention = ".mwf"
	Mxf         FileExtention = ".mxf"
	Mxl         FileExtention = ".mxl"
	Mxml        FileExtention = ".mxml"
	Mxs         FileExtention = ".mxs"
	Mxu         FileExtention = ".mxu"
	N3          FileExtention = ".n3"
	NGage       FileExtention = ".n-gage"
	Nbp         FileExtention = ".nbp"
	Nc          FileExtention = ".nc"
	Ncx         FileExtention = ".ncx"
	Ngdat       FileExtention = ".ngdat"
	Nlu         FileExtention = ".nlu"
	Nml         FileExtention = ".nml"
	Nnd         FileExtention = ".nnd"
	Nns         FileExtention = ".nns"
	Nnw         FileExtention = ".nnw"
	Npx         FileExtention = ".npx"
	Nsf         FileExtention = ".nsf"
	Oa2         FileExtention = ".oa2"
	Oa3         FileExtention = ".oa3"
	Oas         FileExtention = ".oas"
	Obd         FileExtention = ".obd"
	Oda         FileExtention = ".oda"
	Odb         FileExtention = ".odb"
	Odc         FileExtention = ".odc"
	Odf         FileExtention = ".odf"
	Odft        FileExtention = ".odft"
	Odg         FileExtention = ".odg"
	Odi         FileExtention = ".odi"
	Odm         FileExtention = ".odm"
	Odp         FileExtention = ".odp"
	Ods         FileExtention = ".ods"
	Odt         FileExtention = ".odt"
	Oga         FileExtention = ".oga"
	Ogv         FileExtention = ".ogv"
	Ogx         FileExtention = ".ogx"
	Onetoc      FileExtention = ".onetoc"
	Opf         FileExtention = ".opf"
	Opus        FileExtention = ".opus"
	Org         FileExtention = ".org"
	Osf         FileExtention = ".osf"
	Osfpvg      FileExtention = ".osfpvg"
	Otc         FileExtention = ".otc"
	Otf         FileExtention = ".otf"
	Otg         FileExtention = ".otg"
	Oth         FileExtention = ".oth"
	Oti         FileExtention = ".oti"
	Otp         FileExtention = ".otp"
	Ots         FileExtention = ".ots"
	Ott         FileExtention = ".ott"
	Oxt         FileExtention = ".oxt"
	P           FileExtention = ".p"
	P10         FileExtention = ".p10"
	P12         FileExtention = ".p12"
	P7b         FileExtention = ".p7b"
	P7m         FileExtention = ".p7m"
	P7r         FileExtention = ".p7r"
	P7s         FileExtention = ".p7s"
	P8          FileExtention = ".p8"
	Par         FileExtention = ".par"
	Paw         FileExtention = ".paw"
	Pbd         FileExtention = ".pbd"
	Pbm         FileExtention = ".pbm"
	Pcf         FileExtention = ".pcf"
	Pcl         FileExtention = ".pcl"
	Pclxl       FileExtention = ".pclxl"
	Pcurl       FileExtention = ".pcurl"
	Pcx         FileExtention = ".pcx"
	Pdb         FileExtention = ".pdb"
	Pdf         FileExtention = ".pdf"
	Pfa         FileExtention = ".pfa"
	Pfr         FileExtention = ".pfr"
	Pgm         FileExtention = ".pgm"
	Pgn         FileExtention = ".pgn"
	Pgp         FileExtention = ".pgp"
	Php         FileExtention = ".php"
	Pic         FileExtention = ".pic"
	Pjpeg       FileExtention = ".pjpeg"
	Pki         FileExtention = ".pki"
	Pkipath     FileExtention = ".pkipath"
	Plb         FileExtention = ".plb"
	Plc         FileExtention = ".plc"
	Plf         FileExtention = ".plf"
	Pls         FileExtention = ".pls"
	Pml         FileExtention = ".pml"
	Png         FileExtention = ".png"
	Pnm         FileExtention = ".pnm"
	Portpkg     FileExtention = ".portpkg"
	Potm        FileExtention = ".potm"
	Potx        FileExtention = ".potx"
	Ppam        FileExtention = ".ppam"
	Ppd         FileExtention = ".ppd"
	Ppm         FileExtention = ".ppm"
	Ppsm        FileExtention = ".ppsm"
	Ppsx        FileExtention = ".ppsx"
	Ppt         FileExtention = ".ppt"
	Pptm        FileExtention = ".pptm"
	Pptx        FileExtention = ".pptx"
	Prc         FileExtention = ".prc"
	Pre         FileExtention = ".pre"
	Prf         FileExtention = ".prf"
	Psb         FileExtention = ".psb"
	Psd         FileExtention = ".psd"
	Psf         FileExtention = ".psf"
	Pskcxml     FileExtention = ".pskcxml"
	Ptid        FileExtention = ".ptid"
	Pub         FileExtention = ".pub"
	Pvb         FileExtention = ".pvb"
	Pwn         FileExtention = ".pwn"
	Pya         FileExtention = ".pya"
	Pyv         FileExtention = ".pyv"
	Qam         FileExtention = ".qam"
	Qbo         FileExtention = ".qbo"
	Qfx         FileExtention = ".qfx"
	Qps         FileExtention = ".qps"
	Qt          FileExtention = ".qt"
	Qxd         FileExtention = ".qxd"
	Ram         FileExtention = ".ram"
	Rar         FileExtention = ".rar"
	Ras         FileExtention = ".ras"
	Rcprofile   FileExtention = ".rcprofile"
	Rdf         FileExtention = ".rdf"
	Rdz         FileExtention = ".rdz"
	Rep         FileExtention = ".rep"
	Res         FileExtention = ".res"
	Rgb         FileExtention = ".rgb"
	Rif         FileExtention = ".rif"
	Rip         FileExtention = ".rip"
	Rl          FileExtention = ".rl"
	Rlc         FileExtention = ".rlc"
	Rld         FileExtention = ".rld"
	Rm          FileExtention = ".rm"
	Rmp         FileExtention = ".rmp"
	Rms         FileExtention = ".rms"
	Rnc         FileExtention = ".rnc"
	Rp9         FileExtention = ".rp9"
	Rpss        FileExtention = ".rpss"
	Rpst        FileExtention = ".rpst"
	Rq          FileExtention = ".rq"
	Rs          FileExtention = ".rs"
	Rsd         FileExtention = ".rsd"
	Rss         FileExtention = ".rss"
	Rtf         FileExtention = ".rtf"
	Rtx         FileExtention = ".rtx"
	S           FileExtention = ".s"
	Saf         FileExtention = ".saf"
	Sbml        FileExtention = ".sbml"
	Sc          FileExtention = ".sc"
	Scd         FileExtention = ".scd"
	Scm         FileExtention = ".scm"
	Scq         FileExtention = ".scq"
	Scs         FileExtention = ".scs"
	Scurl       FileExtention = ".scurl"
	Sda         FileExtention = ".sda"
	Sdc         FileExtention = ".sdc"
	Sdd         FileExtention = ".sdd"
	Sdkm        FileExtention = ".sdkm"
	Sdp         FileExtention = ".sdp"
	Sdw         FileExtention = ".sdw"
	See         FileExtention = ".see"
	Seed        FileExtention = ".seed"
	Sema        FileExtention = ".sema"
	Semd        FileExtention = ".semd"
	Semf        FileExtention = ".semf"
	Ser         FileExtention = ".ser"
	Setpay      FileExtention = ".setpay"
	Setreg      FileExtention = ".setreg"
	SfdHdstx    FileExtention = ".sfd-hdstx"
	Sfs         FileExtention = ".sfs"
	Sgl         FileExtention = ".sgl"
	Sgml        FileExtention = ".sgml"
	Sh          FileExtention = ".sh"
	Shar        FileExtention = ".shar"
	Shf         FileExtention = ".shf"
	Sis         FileExtention = ".sis"
	Sit         FileExtention = ".sit"
	Sitx        FileExtention = ".sitx"
	Skp         FileExtention = ".skp"
	Sldm        FileExtention = ".sldm"
	Sldx        FileExtention = ".sldx"
	Slt         FileExtention = ".slt"
	Sm          FileExtention = ".sm"
	Smf         FileExtention = ".smf"
	Smi         FileExtention = ".smi"
	Snf         FileExtention = ".snf"
	Spf         FileExtention = ".spf"
	Spl         FileExtention = ".spl"
	Spot        FileExtention = ".spot"
	Spp         FileExtention = ".spp"
	Spq         FileExtention = ".spq"
	Src         FileExtention = ".src"
	Sru         FileExtention = ".sru"
	Srx         FileExtention = ".srx"
	Sse         FileExtention = ".sse"
	Ssf         FileExtention = ".ssf"
	Ssml        FileExtention = ".ssml"
	St          FileExtention = ".st"
	Stc         FileExtention = ".stc"
	Std         FileExtention = ".std"
	Stf         FileExtention = ".stf"
	Sti         FileExtention = ".sti"
	Stk         FileExtention = ".stk"
	Stl         FileExtention = ".stl"
	Str         FileExtention = ".str"
	Stw         FileExtention = ".stw"
	Sub         FileExtention = ".sub"
	Sus         FileExtention = ".sus"
	Sv4cpio     FileExtention = ".sv4cpio"
	Sv4crc      FileExtention = ".sv4crc"
	Svc         FileExtention = ".svc"
	Svd         FileExtention = ".svd"
	Svg         FileExtention = ".svg"
	Swf         FileExtention = ".swf"
	Swi         FileExtention = ".swi"
	Sxc         FileExtention = ".sxc"
	Sxd         FileExtention = ".sxd"
	Sxg         FileExtention = ".sxg"
	Sxi         FileExtention = ".sxi"
	Sxm         FileExtention = ".sxm"
	Sxw         FileExtention = ".sxw"
	T           FileExtention = ".t"
	Tao         FileExtention = ".tao"
	Tar         FileExtention = ".tar"
	Tcap        FileExtention = ".tcap"
	Tcl         FileExtention = ".tcl"
	Teacher     FileExtention = ".teacher"
	Tei         FileExtention = ".tei"
	Tex         FileExtention = ".tex"
	Texinfo     FileExtention = ".texinfo"
	Tfi         FileExtention = ".tfi"
	Tfm         FileExtention = ".tfm"
	Thmx        FileExtention = ".thmx"
	Tif         FileExtention = ".tif"
	Tiff        FileExtention = ".tiff"
	Tmo         FileExtention = ".tmo"
	Torrent     FileExtention = ".torrent"
	Tpl         FileExtention = ".tpl"
	Tpt         FileExtention = ".tpt"
	Tra         FileExtention = ".tra"
	Trm         FileExtention = ".trm"
	Ts          FileExtention = ".ts"
	Tsd         FileExtention = ".tsd"
	Tsv         FileExtention = ".tsv"
	Ttf         FileExtention = ".ttf"
	Ttl         FileExtention = ".ttl"
	Twd         FileExtention = ".twd"
	Txd         FileExtention = ".txd"
	Txf         FileExtention = ".txf"
	Txt         FileExtention = ".txt"
	Ufd         FileExtention = ".ufd"
	Umj         FileExtention = ".umj"
	Unityweb    FileExtention = ".unityweb"
	Uoml        FileExtention = ".uoml"
	Uri         FileExtention = ".uri"
	Ustar       FileExtention = ".ustar"
	Utz         FileExtention = ".utz"
	Uu          FileExtention = ".uu"
	Uva         FileExtention = ".uva"
	Uvh         FileExtention = ".uvh"
	Uvi         FileExtention = ".uvi"
	Uvm         FileExtention = ".uvm"
	Uvp         FileExtention = ".uvp"
	Uvs         FileExtention = ".uvs"
	Uvu         FileExtention = ".uvu"
	Uvv         FileExtention = ".uvv"
	Vcd         FileExtention = ".vcd"
	Vcf         FileExtention = ".vcf"
	Vcg         FileExtention = ".vcg"
	Vcs         FileExtention = ".vcs"
	Vcx         FileExtention = ".vcx"
	Vis         FileExtention = ".vis"
	Viv         FileExtention = ".viv"
	Vsd         FileExtention = ".vsd"
	Vsdx        FileExtention = ".vsdx"
	Vsf         FileExtention = ".vsf"
	Vtu         FileExtention = ".vtu"
	Vxml        FileExtention = ".vxml"
	Wad         FileExtention = ".wad"
	Wav         FileExtention = ".wav"
	Wax         FileExtention = ".wax"
	Wbmp        FileExtention = ".wbmp"
	Wbs         FileExtention = ".wbs"
	Wbxml       FileExtention = ".wbxml"
	Weba        FileExtention = ".weba"
	Webm        FileExtention = ".webm"
	Webp        FileExtention = ".webp"
	Wg          FileExtention = ".wg"
	Wgt         FileExtention = ".wgt"
	Wm          FileExtention = ".wm"
	Wma         FileExtention = ".wma"
	Wmd         FileExtention = ".wmd"
	Wmf         FileExtention = ".wmf"
	Wml         FileExtention = ".wml"
	Wmlc        FileExtention = ".wmlc"
	Wmls        FileExtention = ".wmls"
	Wmlsc       FileExtention = ".wmlsc"
	Wmv         FileExtention = ".wmv"
	Wmx         FileExtention = ".wmx"
	Wmz         FileExtention = ".wmz"
	Woff        FileExtention = ".woff"
	Woff2       FileExtention = ".woff2"
	Wpd         FileExtention = ".wpd"
	Wpl         FileExtention = ".wpl"
	Wps         FileExtention = ".wps"
	Wqd         FileExtention = ".wqd"
	Wri         FileExtention = ".wri"
	Wrl         FileExtention = ".wrl"
	Wsdl        FileExtention = ".wsdl"
	Wspolicy    FileExtention = ".wspolicy"
	Wtb         FileExtention = ".wtb"
	Wvx         FileExtention = ".wvx"
	X3d         FileExtention = ".x3d"
	Xap         FileExtention = ".xap"
	Xar         FileExtention = ".xar"
	Xbap        FileExtention = ".xbap"
	Xbd         FileExtention = ".xbd"
	Xbm         FileExtention = ".xbm"
	Xdf         FileExtention = ".xdf"
	Xdm         FileExtention = ".xdm"
	Xdp         FileExtention = ".xdp"
	Xdssc       FileExtention = ".xdssc"
	Xdw         FileExtention = ".xdw"
	Xenc        FileExtention = ".xenc"
	Xer         FileExtention = ".xer"
	Xfdf        FileExtention = ".xfdf"
	Xfdl        FileExtention = ".xfdl"
	Xhtml       FileExtention = ".xhtml"
	Xif         FileExtention = ".xif"
	Xlam        FileExtention = ".xlam"
	Xls         FileExtention = ".xls"
	Xlsb        FileExtention = ".xlsb"
	Xlsm        FileExtention = ".xlsm"
	Xlsx        FileExtention = ".xlsx"
	Xltm        FileExtention = ".xltm"
	Xltx        FileExtention = ".xltx"
	Xml         FileExtention = ".xml"
	Xo          FileExtention = ".xo"
	Xop         FileExtention = ".xop"
	Xpi         FileExtention = ".xpi"
	Xpm         FileExtention = ".xpm"
	Xpr         FileExtention = ".xpr"
	Xps         FileExtention = ".xps"
	Xpw         FileExtention = ".xpw"
	Xslt        FileExtention = ".xslt"
	Xsm         FileExtention = ".xsm"
	Xspf        FileExtention = ".xspf"
	Xul         FileExtention = ".xul"
	Xwd         FileExtention = ".xwd"
	Xyz         FileExtention = ".xyz"
	Yaml        FileExtention = ".yaml"
	Yang        FileExtention = ".yang"
	Yin         FileExtention = ".yin"
	Yml         FileExtention = ".yml"
	Zaz         FileExtention = ".zaz"
	Zip         FileExtention = ".zip"
	Zir         FileExtention = ".zir"
	Zmm         FileExtention = ".zmm"
)

var (
	All = FileExtentions{
		_123,
		_3dml,
		_3g2,
		_3gp,
		_7z,
		Aab,
		Aac,
		Aam,
		Aas,
		Abw,
		Ac,
		Acc,
		Ace,
		Acu,
		Adp,
		Aep,
		Afp,
		Ahead,
		Ai,
		Aif,
		Air,
		Ait,
		Ami,
		Apk,
		Application,
		Apr,
		Arc,
		Asf,
		Aso,
		Atc,
		Atom,
		Atomcat,
		Atomsvc,
		Atx,
		Au,
		Avi,
		Aw,
		Azf,
		Azs,
		Azw,
		Bcpio,
		Bdf,
		Bdm,
		Bed,
		Bh2,
		Bin,
		Bmi,
		Bmp,
		Box,
		Btif,
		Bz,
		Bz2,
		C,
		C11amc,
		C11amz,
		C4g,
		Cab,
		Car,
		Cat,
		Ccxml,
		Cda,
		Cdbcmsg,
		Cdkey,
		Cdmia,
		Cdmic,
		Cdmid,
		Cdmio,
		Cdmiq,
		Cdx,
		Cdxml,
		Cdy,
		Cer,
		Cgm,
		Chat,
		Chm,
		Chrt,
		Cif,
		Cii,
		Cil,
		Cla,
		Class,
		Clkk,
		Clkp,
		Clkt,
		Clkw,
		Clkx,
		Clp,
		Cmc,
		Cmdf,
		Cml,
		Cmp,
		Cmx,
		Cod,
		Cpio,
		Cpt,
		Crd,
		Crl,
		Cryptonote,
		Csh,
		Csml,
		Csp,
		Css,
		Csv,
		Cu,
		Curl,
		Cww,
		Dae,
		Daf,
		Davmount,
		Dcurl,
		Dd2,
		Ddd,
		Deb,
		Der,
		Dfac,
		Dir,
		Dis,
		Djvu,
		Dmg,
		Dna,
		Doc,
		Docm,
		Docx,
		Dotm,
		Dotx,
		Dp,
		Dpg,
		Dra,
		Dsc,
		Dssc,
		Dtb,
		Dtd,
		Dts,
		Dtshd,
		Dvi,
		Dwf,
		Dwg,
		Dxf,
		Dxp,
		Ecelp4800,
		Ecelp7470,
		Ecelp9600,
		Edm,
		Edx,
		Efif,
		Ei6,
		Eml,
		Emma,
		Eol,
		Eot,
		Epub,
		Es,
		Es3,
		Esf,
		Etx,
		Exe,
		Exi,
		Ext,
		Ez2,
		Ez3,
		F,
		F4v,
		Fbs,
		Fcs,
		Fdf,
		FeLaunch,
		Fg5,
		Fh,
		Fig,
		Fli,
		Flo,
		Flv,
		Flw,
		Flx,
		Fly,
		Fm,
		Fnc,
		Fpx,
		Fsc,
		Fst,
		Ftc,
		Fti,
		Fvt,
		Fxp,
		Fzs,
		G2w,
		G3,
		G3w,
		Gac,
		Gdl,
		Geo,
		Gex,
		Ggb,
		Ggt,
		Ghf,
		Gif,
		Gim,
		Gmx,
		Gnumeric,
		Gph,
		Gpx,
		Gqf,
		Gram,
		Grv,
		Grxml,
		Gsf,
		Gtar,
		Gtm,
		Gtw,
		Gv,
		Gxt,
		Gz,
		H261,
		H263,
		H264,
		Hal,
		Hbci,
		Hdf,
		Hlp,
		Hpgl,
		Hpid,
		Hps,
		Hqx,
		Htke,
		Html,
		Hvd,
		Hvp,
		Hvs,
		I2g,
		Icc,
		Ice,
		Ico,
		Ics,
		Ief,
		Ifm,
		Igl,
		Igm,
		Igs,
		Igx,
		Iif,
		Imp,
		Ims,
		Ipfix,
		Ipk,
		Irm,
		Irp,
		Itp,
		Ivp,
		Ivu,
		Jad,
		Jam,
		Jar,
		Java,
		Jisp,
		Jlt,
		Jnlp,
		Joda,
		Jpeg,
		Jpg,
		Jpgv,
		Jpm,
		Js,
		Json,
		Jsonld,
		Karbon,
		Kfo,
		Kia,
		Kml,
		Kmz,
		Kne,
		Kon,
		Kpr,
		Ksp,
		Ktx,
		Ktz,
		Kwd,
		Lasxml,
		Latex,
		Lbd,
		Lbe,
		Les,
		Link66,
		Lrm,
		Ltf,
		Lvp,
		Lwp,
		M21,
		M3u,
		M3u8,
		M4v,
		Ma,
		Mads,
		Mag,
		Mathml,
		Mbk,
		Mbox,
		Mc1,
		Mcd,
		Mcurl,
		Mdb,
		Mdi,
		Meta4,
		Mets,
		Mfm,
		Mgp,
		Mgz,
		Mid,
		Midi,
		Mif,
		Mj2,
		Mlp,
		Mmd,
		Mmf,
		Mmr,
		Mny,
		Mods,
		Movie,
		Mp4,
		Mp4a,
		Mpc,
		Mpeg,
		Mpga,
		Mpkg,
		Mpm,
		Mpn,
		Mpp,
		Mpy,
		Mqy,
		Mrc,
		Mrcx,
		Mscml,
		Mseq,
		Msf,
		Msh,
		Msl,
		Msty,
		Mts,
		Mus,
		Musicxml,
		Mvb,
		Mwf,
		Mxf,
		Mxl,
		Mxml,
		Mxs,
		Mxu,
		N3,
		NGage,
		Nbp,
		Nc,
		Ncx,
		Ngdat,
		Nlu,
		Nml,
		Nnd,
		Nns,
		Nnw,
		Npx,
		Nsf,
		Oa2,
		Oa3,
		Oas,
		Obd,
		Oda,
		Odb,
		Odc,
		Odf,
		Odft,
		Odg,
		Odi,
		Odm,
		Odp,
		Ods,
		Odt,
		Oga,
		Ogv,
		Ogx,
		Onetoc,
		Opf,
		Opus,
		Org,
		Osf,
		Osfpvg,
		Otc,
		Otf,
		Otg,
		Oth,
		Oti,
		Otp,
		Ots,
		Ott,
		Oxt,
		P,
		P10,
		P12,
		P7b,
		P7m,
		P7r,
		P7s,
		P8,
		Par,
		Paw,
		Pbd,
		Pbm,
		Pcf,
		Pcl,
		Pclxl,
		Pcurl,
		Pcx,
		Pdb,
		Pdf,
		Pfa,
		Pfr,
		Pgm,
		Pgn,
		Pgp,
		Php,
		Pic,
		Pjpeg,
		Pki,
		Pkipath,
		Plb,
		Plc,
		Plf,
		Pls,
		Pml,
		Png,
		Pnm,
		Portpkg,
		Potm,
		Potx,
		Ppam,
		Ppd,
		Ppm,
		Ppsm,
		Ppsx,
		Ppt,
		Pptm,
		Pptx,
		Prc,
		Pre,
		Prf,
		Psb,
		Psd,
		Psf,
		Pskcxml,
		Ptid,
		Pub,
		Pvb,
		Pwn,
		Pya,
		Pyv,
		Qam,
		Qbo,
		Qfx,
		Qps,
		Qt,
		Qxd,
		Ram,
		Rar,
		Ras,
		Rcprofile,
		Rdf,
		Rdz,
		Rep,
		Res,
		Rgb,
		Rif,
		Rip,
		Rl,
		Rlc,
		Rld,
		Rm,
		Rmp,
		Rms,
		Rnc,
		Rp9,
		Rpss,
		Rpst,
		Rq,
		Rs,
		Rsd,
		Rss,
		Rtf,
		Rtx,
		S,
		Saf,
		Sbml,
		Sc,
		Scd,
		Scm,
		Scq,
		Scs,
		Scurl,
		Sda,
		Sdc,
		Sdd,
		Sdkm,
		Sdp,
		Sdw,
		See,
		Seed,
		Sema,
		Semd,
		Semf,
		Ser,
		Setpay,
		Setreg,
		SfdHdstx,
		Sfs,
		Sgl,
		Sgml,
		Sh,
		Shar,
		Shf,
		Sis,
		Sit,
		Sitx,
		Skp,
		Sldm,
		Sldx,
		Slt,
		Sm,
		Smf,
		Smi,
		Snf,
		Spf,
		Spl,
		Spot,
		Spp,
		Spq,
		Src,
		Sru,
		Srx,
		Sse,
		Ssf,
		Ssml,
		St,
		Stc,
		Std,
		Stf,
		Sti,
		Stk,
		Stl,
		Str,
		Stw,
		Sub,
		Sus,
		Sv4cpio,
		Sv4crc,
		Svc,
		Svd,
		Svg,
		Swf,
		Swi,
		Sxc,
		Sxd,
		Sxg,
		Sxi,
		Sxm,
		Sxw,
		T,
		Tao,
		Tar,
		Tcap,
		Tcl,
		Teacher,
		Tei,
		Tex,
		Texinfo,
		Tfi,
		Tfm,
		Thmx,
		Tif,
		Tiff,
		Tmo,
		Torrent,
		Tpl,
		Tpt,
		Tra,
		Trm,
		Ts,
		Tsd,
		Tsv,
		Ttf,
		Ttl,
		Twd,
		Txd,
		Txf,
		Txt,
		Ufd,
		Umj,
		Unityweb,
		Uoml,
		Uri,
		Ustar,
		Utz,
		Uu,
		Uva,
		Uvh,
		Uvi,
		Uvm,
		Uvp,
		Uvs,
		Uvu,
		Uvv,
		Vcd,
		Vcf,
		Vcg,
		Vcs,
		Vcx,
		Vis,
		Viv,
		Vsd,
		Vsdx,
		Vsf,
		Vtu,
		Vxml,
		Wad,
		Wav,
		Wax,
		Wbmp,
		Wbs,
		Wbxml,
		Weba,
		Webm,
		Webp,
		Wg,
		Wgt,
		Wm,
		Wma,
		Wmd,
		Wmf,
		Wml,
		Wmlc,
		Wmls,
		Wmlsc,
		Wmv,
		Wmx,
		Wmz,
		Woff,
		Woff2,
		Wpd,
		Wpl,
		Wps,
		Wqd,
		Wri,
		Wrl,
		Wsdl,
		Wspolicy,
		Wtb,
		Wvx,
		X3d,
		Xap,
		Xar,
		Xbap,
		Xbd,
		Xbm,
		Xdf,
		Xdm,
		Xdp,
		Xdssc,
		Xdw,
		Xenc,
		Xer,
		Xfdf,
		Xfdl,
		Xhtml,
		Xif,
		Xlam,
		Xls,
		Xlsb,
		Xlsm,
		Xlsx,
		Xltm,
		Xltx,
		Xml,
		Xo,
		Xop,
		Xpi,
		Xpm,
		Xpr,
		Xps,
		Xpw,
		Xslt,
		Xsm,
		Xspf,
		Xul,
		Xwd,
		Xyz,
		Yaml,
		Yang,
		Yin,
		Yml,
		Zaz,
		Zip,
		Zir,
		Zmm,
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) (FileExtention, error) {
	return All.Parse(v)
}

func Strings() []string {
	return All.Strings()
}

func (t FileExtention) String() string {
	return string(t)
}

var ErrFileExtentionInvalid = errors.New("invalid enumeration type")

func Error(vs FileExtentions, v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrFileExtentionInvalid, v, strings.Join(vs.Strings(), ","),
	)
}

func (t FileExtentions) Strings() []string {
	var ss []string

	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t FileExtentions) Parse(v string) (FileExtention, error) {
	var o FileExtention
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Error(t, v)
	}

	return o, nil
}

func (t FileExtentions) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t FileExtentions) Contains(v FileExtention) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t FileExtention) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *FileExtention) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}