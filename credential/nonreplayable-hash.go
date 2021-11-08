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

	"github.com/maxlandon/gondor/maltego"

	"github.com/maxlandon/aims/proto/gen/go/credential"
)

// NonReplayableHash - A credential.PasswordHash password hash that cannot be replayed to authenticate
// to other services. Contrasts with credential.ReplayableHash. The NonReplayableHash.Data is any password
// hash, such as those recovered from `/etc/passwd` or `/etc/shadow`.
type NonReplayableHash PasswordHash

// NewNonReplayableHash - Create a new NonReplayableHash Credential.
func NewNonReplayableHash() *NonReplayableHash {
	h := NonReplayableHash(PasswordHash{})
	h.Type = credential.PrivateType_NonReplayableHash
	return &h
}

//
// General Functions
//

// ToORM - Get the SQL object for the NonReplayableHash credential.
func (h *NonReplayableHash) ToORM(ctx context.Context) (credential.PrivateORM, error) {
	return (*PasswordHash)(h).ToORM(ctx)
}

// ToPB - Get the Protobuf object for the NonReplayableHash credential.
func (h *NonReplayableHash) ToPB(ctx context.Context) *credential.Private {
	return (*PasswordHash)(h).ToPB(ctx)
}

// AsEntity - Returns the Private as a valid Maltego Entity.
func (h *NonReplayableHash) AsEntity() maltego.Entity {
	// e:= maltego.NewEntity(h)
	// base := (*Private)(h).AsEntity()
	// e.SetBase(base)
	// return e
	return maltego.NewEntity(h)
}