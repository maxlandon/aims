package scan

/*
   AIMS (Attacked Infrastructure Modular Specification)
   Copyright (C) 2021 Maxime Landon

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"

	"github.com/maxlandon/aims/proto/scan"
)

// Target - This type can be used as an Input object to a scan,
// in which case only the Input fields matter to you
//
// Represents how the target was specified when passed to nmap,
// its status and the reason of its status. Example:
// <target specification="domain.does.not.exist" status="skipped" reason="invalid"/>
type Target scan.Target

// ToORM - Get the SQL object for the Target
func (t *Target) ToORM(ctx context.Context) (scan.TargetORM, error) {
	return (*scan.Target)(t).ToORM(ctx)
}

// ToPB - Get the Protobuf object for the Target
func (t *Target) ToPB() *scan.Target {
	return (*scan.Target)(t)
}
