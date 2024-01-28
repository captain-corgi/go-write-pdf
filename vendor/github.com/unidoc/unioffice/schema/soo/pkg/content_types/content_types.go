//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package content_types ;import (_a "encoding/xml";_ec "fmt";_dd "github.com/unidoc/unioffice";_b "github.com/unidoc/unioffice/common/logger";_d "regexp";);

// Validate validates the CT_Types and its children
func (_cgee *CT_Types )Validate ()error {return _cgee .ValidateWithPath ("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073");};const ST_ExtensionPattern ="\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b";
var ST_ExtensionPatternRe =_d .MustCompile (ST_ExtensionPattern );func NewCT_Override ()*CT_Override {_dg :=&CT_Override {};_dg .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _dg ;};func (_gg *CT_Types )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_ca :for {_egb ,_bf :=d .Token ();
if _bf !=nil {return _bf ;};switch _ab :=_egb .(type ){case _a .StartElement :switch _ab .Name {case _a .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_cdg :=NewDefault ();
if _dfe :=d .DecodeElement (_cdg ,&_ab );_dfe !=nil {return _dfe ;};_gg .Default =append (_gg .Default ,_cdg );case _a .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_ed :=NewOverride ();
if _agg :=d .DecodeElement (_ed ,&_ab );_agg !=nil {return _agg ;};_gg .Override =append (_gg .Override ,_ed );default:_b .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076",_ab .Name );
if _ddg :=d .Skip ();_ddg !=nil {return _ddg ;};};case _a .EndElement :break _ca ;case _a .CharData :};};return nil ;};func (_fg *Override )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {return _fg .CT_Override .MarshalXML (e ,start );};

// Validate validates the Override and its children
func (_fdg *Override )Validate ()error {return _fdg .ValidateWithPath ("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};func (_ac *CT_Default )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"},Value :_ec .Sprintf ("\u0025\u0076",_ac .ExtensionAttr )});
start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_ec .Sprintf ("\u0025\u0076",_ac .ContentTypeAttr )});e .EncodeToken (start );e .EncodeToken (_a .EndElement {Name :start .Name });
return nil ;};type Override struct{CT_Override };func (_feca *Override )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_feca .CT_Override =*NewCT_Override ();for _ ,_ecc :=range start .Attr {if _ecc .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_efe ,_bad :=_ecc .Value ,error (nil );
if _bad !=nil {return _bad ;};_feca .ContentTypeAttr =_efe ;continue ;};if _ecc .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_bcg ,_fag :=_ecc .Value ,error (nil );if _fag !=nil {return _fag ;};_feca .PartNameAttr =_bcg ;continue ;
};};for {_bfc ,_dgf :=d .Token ();if _dgf !=nil {return _ec .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073",_dgf );};if _ce ,_gfg :=_bfc .(_a .EndElement );_gfg &&_ce .Name ==start .Name {break ;
};};return nil ;};func NewCT_Default ()*CT_Default {_f :=&CT_Default {};_f .ExtensionAttr ="\u0078\u006d\u006c";_f .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _f ;};type CT_Types struct{Default []*Default ;
Override []*Override ;};func (_da *CT_Types )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {e .EncodeToken (start );if _da .Default !=nil {_fa :=_a .StartElement {Name :_a .Name {Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}};for _ ,_cge :=range _da .Default {e .EncodeElement (_cge ,_fa );
};};if _da .Override !=nil {_adc :=_a .StartElement {Name :_a .Name {Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}};for _ ,_gd :=range _da .Override {e .EncodeElement (_gd ,_adc );};};e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;
};func (_gf *Default )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_gf .CT_Default =*NewCT_Default ();for _ ,_eeg :=range start .Attr {if _eeg .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_dff ,_fd :=_eeg .Value ,error (nil );
if _fd !=nil {return _fd ;};_gf .ExtensionAttr =_dff ;continue ;};if _eeg .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_egc ,_ece :=_eeg .Value ,error (nil );if _ece !=nil {return _ece ;};_gf .ContentTypeAttr =_egc ;continue ;
};};for {_bc ,_ega :=d .Token ();if _ega !=nil {return _ec .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073",_ega );};if _gfb ,_ggc :=_bc .(_a .EndElement );_ggc &&_gfb .Name ==start .Name {break ;
};};return nil ;};

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_aaa *CT_Types )ValidateWithPath (path string )error {for _gaa ,_fc :=range _aaa .Default {if _dcb :=_fc .ValidateWithPath (_ec .Sprintf ("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d",path ,_gaa ));_dcb !=nil {return _dcb ;
};};for _daf ,_fec :=range _aaa .Override {if _dcc :=_fec .ValidateWithPath (_ec .Sprintf ("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d",path ,_daf ));_dcc !=nil {return _dcc ;};};return nil ;};func NewCT_Types ()*CT_Types {_cda :=&CT_Types {};
return _cda };type CT_Default struct{ExtensionAttr string ;ContentTypeAttr string ;};const ST_ContentTypePattern ="\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024";type Default struct{CT_Default };type CT_Override struct{ContentTypeAttr string ;
PartNameAttr string ;};func (_eac *Types )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_eac .CT_Types =*NewCT_Types ();_eda :for {_cce ,_eca :=d .Token ();if _eca !=nil {return _eca ;};switch _ecag :=_cce .(type ){case _a .StartElement :switch _ecag .Name {case _a .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_gge :=NewDefault ();
if _cdb :=d .DecodeElement (_gge ,&_ecag );_cdb !=nil {return _cdb ;};_eac .Default =append (_eac .Default ,_gge );case _a .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_adcd :=NewOverride ();
if _aaf :=d .DecodeElement (_adcd ,&_ecag );_aaf !=nil {return _aaf ;};_eac .Override =append (_eac .Override ,_adcd );default:_b .Log .Debug ("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076",_ecag .Name );
if _abb :=d .Skip ();_abb !=nil {return _abb ;};};case _a .EndElement :break _eda ;case _a .CharData :};};return nil ;};func (_c *CT_Default )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_c .ExtensionAttr ="\u0078\u006d\u006c";_c .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";
for _ ,_db :=range start .Attr {if _db .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_cf ,_ba :=_db .Value ,error (nil );if _ba !=nil {return _ba ;};_c .ExtensionAttr =_cf ;continue ;};if _db .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_af ,_fe :=_db .Value ,error (nil );
if _fe !=nil {return _fe ;};_c .ContentTypeAttr =_af ;continue ;};};for {_ad ,_g :=d .Token ();if _g !=nil {return _ec .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073",_g );
};if _ef ,_ee :=_ad .(_a .EndElement );_ee &&_ef .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_dgc *Types )ValidateWithPath (path string )error {if _eb :=_dgc .CT_Types .ValidateWithPath (path );_eb !=nil {return _eb ;};return nil ;};func NewOverride ()*Override {_bga :=&Override {};_bga .CT_Override =*NewCT_Override ();return _bga };

// Validate validates the Types and its children
func (_ddc *Types )Validate ()error {return _ddc .ValidateWithPath ("\u0054\u0079\u0070e\u0073")};

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_afa *CT_Default )ValidateWithPath (path string )error {if !ST_ExtensionPatternRe .MatchString (_afa .ExtensionAttr ){return _ec .Errorf ("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029",path ,ST_ExtensionPatternRe ,_afa .ExtensionAttr );
};if !ST_ContentTypePatternRe .MatchString (_afa .ContentTypeAttr ){return _ec .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_afa .ContentTypeAttr );
};return nil ;};type Types struct{CT_Types };var ST_ContentTypePatternRe =_d .MustCompile (ST_ContentTypePattern );

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_bg *Default )ValidateWithPath (path string )error {if _ggg :=_bg .CT_Default .ValidateWithPath (path );_ggg !=nil {return _ggg ;};return nil ;};func NewDefault ()*Default {_fb :=&Default {};_fb .CT_Default =*NewCT_Default ();return _fb };

// Validate validates the CT_Override and its children
func (_be *CT_Override )Validate ()error {return _be .ValidateWithPath ("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};func (_ga *CT_Override )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_ec .Sprintf ("\u0025\u0076",_ga .ContentTypeAttr )});
start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"},Value :_ec .Sprintf ("\u0025\u0076",_ga .PartNameAttr )});e .EncodeToken (start );e .EncodeToken (_a .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_bge *Override )ValidateWithPath (path string )error {if _efb :=_bge .CT_Override .ValidateWithPath (path );_efb !=nil {return _efb ;};return nil ;};func (_ea *Default )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {return _ea .CT_Default .MarshalXML (e ,start );
};func NewTypes ()*Types {_cc :=&Types {};_cc .CT_Types =*NewCT_Types ();return _cc };

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_cfc *CT_Override )ValidateWithPath (path string )error {if !ST_ContentTypePatternRe .MatchString (_cfc .ContentTypeAttr ){return _ec .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_cfc .ContentTypeAttr );
};return nil ;};func (_aca *Types )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"});
start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0054\u0079\u0070e\u0073";return _aca .CT_Types .MarshalXML (e ,start );};

// Validate validates the CT_Default and its children
func (_ag *CT_Default )Validate ()error {return _ag .ValidateWithPath ("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074");};

// Validate validates the Default and its children
func (_dgge *Default )Validate ()error {return _dgge .ValidateWithPath ("\u0044e\u0066\u0061\u0075\u006c\u0074");};func (_bb *CT_Override )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_bb .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";
for _ ,_bd :=range start .Attr {if _bd .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_ff ,_dgg :=_bd .Value ,error (nil );if _dgg !=nil {return _dgg ;};_bb .ContentTypeAttr =_ff ;continue ;};if _bd .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_cd ,_dc :=_bd .Value ,error (nil );
if _dc !=nil {return _dc ;};_bb .PartNameAttr =_cd ;continue ;};};for {_eg ,_dbf :=d .Token ();if _dbf !=nil {return _ec .Errorf ("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073",_dbf );
};if _baf ,_cg :=_eg .(_a .EndElement );_cg &&_baf .Name ==start .Name {break ;};};return nil ;};func init (){_dd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073",NewCT_Types );
_dd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074",NewCT_Default );
_dd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewCT_Override );
_dd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0054\u0079\u0070e\u0073",NewTypes );
_dd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0044e\u0066\u0061\u0075\u006c\u0074",NewDefault );
_dd .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewOverride );
};