/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plugins

import (
	pkgCConfigs "github.com/nitroci/nitroci-core/pkg/core/configs"
	pkgCContexts "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"
)

func OnConfigure(context *pkgCContexts.RuntimeContext, args []string, fields map[string]interface{}) {
	// Persist the domain
	var domain string
	if fields["bitbucket-workspace"] != nil {
		domain = fields["bitbucket-workspace"].(string)
	}
	if len(domain) == 0 {
		domain, _ = pkgCTerminal.PromptGlobalConfigKey(context.Cli.Profile, "Workspace", false)
	}
	pkgCConfigs.SetGlobalConfigString(context.Cli.Profile, "bitbucket_workspace", domain)
	// Persist the username
	var username string
	if fields["bitbucket-workspace"] != nil {
		username = fields["bitbucket-user"].(string)
	}
	if len(username) == 0 {
		username, _ = pkgCTerminal.PromptGlobalConfigKey(context.Cli.Profile, "Username", false)
	}
	pkgCConfigs.SetGlobalConfigString(context.Cli.Profile, "bitbucket_username", username)
	// Persist application password
	var password string
	if fields["bitbucket-pass"] != nil {
		password = fields["bitbucket-pass"].(string)
	}
	if len(password) == 0 {
		password, _ = pkgCTerminal.PromptGlobalConfigKey(context.Cli.Profile, "Password", true)
		pkgCConfigs.SetGlobalConfigString(context.Cli.Profile, "bitbucket_secret", password)
	}
	pkgCConfigs.SetGlobalConfigString(context.Cli.Profile, "bitbucket_secret", password)
}
