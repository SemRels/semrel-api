// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The SemRels Authors

// Package plugin provides the go-plugin integration for semrel plugins.
// It defines the GRPCPlugin implementations for each plugin type so that
// both the host (semrel) and plugin binaries share the same handshake config
// and transport wiring.
package plugin

import (
	"context"

	semrelv1 "github.com/GoSemantics/go-semrel-api/api/gen/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// HandshakeConfig is shared between host and all plugin binaries.
// The ProtocolVersion must be bumped when the gRPC interface changes in a
// backwards-incompatible way.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SEMREL_PLUGIN",
	MagicCookieValue: "semrel-v1",
}

// PluginMap lists every supported plugin type.  Hosts and plugin binaries must
// use the same map key for the plugin type they handle.
var PluginMap = map[string]plugin.Plugin{
	"provider":  &ProviderGRPCPlugin{},
	"condition": &CIConditionGRPCPlugin{},
	"analyzer":  &CommitAnalyzerGRPCPlugin{},
	"generator": &ChangelogGeneratorGRPCPlugin{},
	"updater":   &FilesUpdaterGRPCPlugin{},
	"hooks":     &HooksGRPCPlugin{},
}

// ---------------------------------------------------------------------------
// Provider plugin
// ---------------------------------------------------------------------------

// ProviderGRPCPlugin wires the ProviderPlugin gRPC service into go-plugin.
type ProviderGRPCPlugin struct {
	plugin.Plugin
	Impl semrelv1.ProviderPluginServer
}

func (p *ProviderGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	semrelv1.RegisterProviderPluginServer(s, p.Impl)
	return nil
}

func (p *ProviderGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return semrelv1.NewProviderPluginClient(c), nil
}

// ---------------------------------------------------------------------------
// CICondition plugin
// ---------------------------------------------------------------------------

// CIConditionGRPCPlugin wires the CIConditionPlugin gRPC service into go-plugin.
type CIConditionGRPCPlugin struct {
	plugin.Plugin
	Impl semrelv1.CIConditionPluginServer
}

func (p *CIConditionGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	semrelv1.RegisterCIConditionPluginServer(s, p.Impl)
	return nil
}

func (p *CIConditionGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return semrelv1.NewCIConditionPluginClient(c), nil
}

// ---------------------------------------------------------------------------
// CommitAnalyzer plugin
// ---------------------------------------------------------------------------

// CommitAnalyzerGRPCPlugin wires the CommitAnalyzerPlugin gRPC service into go-plugin.
type CommitAnalyzerGRPCPlugin struct {
	plugin.Plugin
	Impl semrelv1.CommitAnalyzerPluginServer
}

func (p *CommitAnalyzerGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	semrelv1.RegisterCommitAnalyzerPluginServer(s, p.Impl)
	return nil
}

func (p *CommitAnalyzerGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return semrelv1.NewCommitAnalyzerPluginClient(c), nil
}

// ---------------------------------------------------------------------------
// ChangelogGenerator plugin
// ---------------------------------------------------------------------------

// ChangelogGeneratorGRPCPlugin wires the ChangelogGeneratorPlugin gRPC service into go-plugin.
type ChangelogGeneratorGRPCPlugin struct {
	plugin.Plugin
	Impl semrelv1.ChangelogGeneratorPluginServer
}

func (p *ChangelogGeneratorGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	semrelv1.RegisterChangelogGeneratorPluginServer(s, p.Impl)
	return nil
}

func (p *ChangelogGeneratorGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return semrelv1.NewChangelogGeneratorPluginClient(c), nil
}

// ---------------------------------------------------------------------------
// FilesUpdater plugin
// ---------------------------------------------------------------------------

// FilesUpdaterGRPCPlugin wires the FilesUpdaterPlugin gRPC service into go-plugin.
type FilesUpdaterGRPCPlugin struct {
	plugin.Plugin
	Impl semrelv1.FilesUpdaterPluginServer
}

func (p *FilesUpdaterGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	semrelv1.RegisterFilesUpdaterPluginServer(s, p.Impl)
	return nil
}

func (p *FilesUpdaterGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return semrelv1.NewFilesUpdaterPluginClient(c), nil
}

// ---------------------------------------------------------------------------
// Hooks plugin
// ---------------------------------------------------------------------------

// HooksGRPCPlugin wires the HooksPlugin gRPC service into go-plugin.
type HooksGRPCPlugin struct {
	plugin.Plugin
	Impl semrelv1.HooksPluginServer
}

func (p *HooksGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	semrelv1.RegisterHooksPluginServer(s, p.Impl)
	return nil
}

func (p *HooksGRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return semrelv1.NewHooksPluginClient(c), nil
}
