//Package configf comment
// This file war generated by tars2go 1.1
// Generated from Config.tars
package configf

import (
	"fmt"
	"github.com/chencaixiong/TarsGo/tars/protocol/codec"
)

//ConfigInfo strcut implement
type ConfigInfo struct {
	Appname       string `json:"appname"`
	Servername    string `json:"servername"`
	Filename      string `json:"filename"`
	BAppOnly      bool   `json:"bAppOnly"`
	Host          string `json:"host"`
	Setdivision   string `json:"setdivision"`
	Containername string `json:"containername"`
}

func (st *ConfigInfo) resetDefault() {
	st.BAppOnly = false
	st.Containername = ""
}

//ReadFrom reads  from _is and put into struct.
func (st *ConfigInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Appname, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Servername, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Filename, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BAppOnly, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Host, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Setdivision, 5, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Containername, 6, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ConfigInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ConfigInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ConfigInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Appname, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Servername, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Filename, 2)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BAppOnly, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Host, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Setdivision, 5)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Containername, 6)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *ConfigInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
