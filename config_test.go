/**
 * @Time : 2022/7/8 6:41 PM
 * @Author : solacowa@gmail.com
 * @File : config_test
 * @Software: GoLand
 */

package config

import "testing"

func TestConfig_GetString(t *testing.T) {
	cfg, err := NewConfig("./app.cfg")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cfg.GetString("server", "app.name"))
}
