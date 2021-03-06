/*
 * Teni-IME - A Vietnamese Input method editor
 * Copyright (C) 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com>
 * This file is part of Teni-IME.
 *
 *  Teni-IME is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Teni-IME is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Teni-IME.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"github.com/sarim/goibus/ibus"
	"os"
	"path/filepath"
)

const (
	DebugComponentName = ComponentName + "Debug"
	DebugEngineName    = EngineName + "-debug"
	IconFile           = "icon.png"
)

func makeDebugComponent() *ibus.Component {

	component := &ibus.Component{
		Name:          "IBusComponent",
		ComponentName: DebugComponentName,
	}

	engine := &ibus.EngineDesc{
		Name:       "IBusEngineDesc",
		EngineName: DebugEngineName,
		Icon:       filepath.Join(filepath.Dir(os.Args[0]), IconFile),
	}

	component.AddEngine(engine)

	return component
}
