package main

import (
	"log"

	runner "github.com/slidebolt/sdk-runner"
	"github.com/slidebolt/sdk-types"
)

type CleanPlugin struct{}

func (p *CleanPlugin) OnInitialize(config runner.Config, state types.Storage) (types.Manifest, types.Storage) {
	return types.Manifest{ID: "plugin-test-clean", Name: "Clean Reference Plugin", Version: "1.0.0"}, state
}

func (p *CleanPlugin) OnReady()                       {}
func (p *CleanPlugin) OnShutdown()                    {}
func (p *CleanPlugin) OnHealthCheck() (string, error) { return "perfect", nil }
func (p *CleanPlugin) OnStorageUpdate(current types.Storage) (types.Storage, error) {
	return current, nil
}

func (p *CleanPlugin) OnDeviceCreate(dev types.Device) (types.Device, error) {
	return dev, nil
}
func (p *CleanPlugin) OnDeviceUpdate(dev types.Device) (types.Device, error) { return dev, nil }
func (p *CleanPlugin) OnDeviceDelete(id string) error                        { return nil }
func (p *CleanPlugin) OnDevicesList(current []types.Device) ([]types.Device, error) {
	return current, nil
}
func (p *CleanPlugin) OnDeviceSearch(q types.SearchQuery, res []types.Device) ([]types.Device, error) {
	return res, nil
}

func (p *CleanPlugin) OnEntityCreate(e types.Entity) (types.Entity, error) { return e, nil }
func (p *CleanPlugin) OnEntityUpdate(e types.Entity) (types.Entity, error) { return e, nil }
func (p *CleanPlugin) OnEntityDelete(d, e string) error                    { return nil }
func (p *CleanPlugin) OnEntitiesList(d string, c []types.Entity) ([]types.Entity, error) {
	return c, nil
}

func (p *CleanPlugin) OnCommand(cmd types.Command, entity types.Entity) (types.Entity, error) {
	return entity, nil
}
func (p *CleanPlugin) OnEvent(evt types.Event, entity types.Entity) (types.Entity, error) {
	return entity, nil
}

func main() {
	if err := runner.NewRunner(&CleanPlugin{}).Run(); err != nil {
		log.Fatal(err)
	}
}
