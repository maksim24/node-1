/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package traversal

import (
	"github.com/mysteriumnetwork/node/services"
)

// NoopPinger does nothing
type NoopPinger struct{}

// Start does nothing
func (np *NoopPinger) Start() {}

// Stop does nothing
func (np *NoopPinger) Stop() {}

// PingProvider does nothing
func (np *NoopPinger) PingProvider(ip string, port int) error { return nil }

// PingTarget does nothing
func (np *NoopPinger) PingTarget(*Params) {}

// BindConsumerPort does nothing
func (np *NoopPinger) BindConsumerPort(port int) {}

// BindServicePort does nothing
func (np *NoopPinger) BindServicePort(serviceType services.ServiceType, port int) {}
