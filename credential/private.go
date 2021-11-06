package credential

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

	"github.com/maxlandon/aims/proto/gen/go/credential"
	"github.com/maxlandon/gondor/maltego"
)

// Private - Base type for all private credentials. A private credential is any credential
// that should not be publicly disclosed, such as a credential.Private.Password password,
// password hash, or key file.
type Private credential.Private

// ToORM - Get the SQL object for the Private credential.
func (p *Private) ToORM(ctx context.Context) (credential.PrivateORM, error) {
	pb := &credential.Private{}
	return (*pb).ToORM(ctx)
}

// AsEntity - Returns the Private as a valid Maltego Entity.
func (h *Private) AsEntity() maltego.Entity {

	return maltego.Entity{}
}
