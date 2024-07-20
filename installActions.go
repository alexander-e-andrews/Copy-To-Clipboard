package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

const REGISTRY_CREATE_NAME = `SOFTWARE\Classes\*\shell\Copy To Clipboard\command`
const REGISTRY_DELETE_KEY_PATH = `SOFTWARE\Classes\*\shell`
const REGISTRY_DELETE_NAME_SUB_PATH = `Copy To Clipboard\command`
const REGISTRY_DELETE_NAME_PATH = `Copy To Clipboard`
// Actions to add copy to clipboard and remove it from the windows system

// Add to registry edit where we are right now
func install(exeLocation string) {
	// Find the registry key
	key, err := getRegistryKey()
	if err != nil {
		return
	}
	defer key.Close()

	// Change its value
	command := fmt.Sprintf("%s \"%%1\"", exeLocation)
	err = key.SetStringValue("", command)
	if err != nil {
		logMessage(fmt.Sprintf("Error accessing registry %v", err))
		return
	}
}

// Remove our data from the registry
// Not going to self delete, leave that up to the user
func uninstall() {
	//Find the registry key
	key, err := getRegistryDeleteKey()
	if err != nil {
		return
	}
	defer key.Close()

	// If we fail to remove our registry entry, that is not good
	// There is probably some kind of rights access that allows you to do this in one action, but I did not see it
	err = registry.DeleteKey(key, REGISTRY_DELETE_NAME_SUB_PATH)
	if err != nil {
		logMessage(fmt.Sprintf("Error deleting registry %v", err))
		return
	}
	err = registry.DeleteKey(key, REGISTRY_DELETE_NAME_PATH)
	if err != nil {
		logMessage(fmt.Sprintf("Error deleting registry %v", err))
		return
	}
}

func getRegistryKey() (key registry.Key, err error) {
	key, _, err = registry.CreateKey(registry.CURRENT_USER, REGISTRY_CREATE_NAME, registry.ALL_ACCESS)
	if err != nil {
		logMessage(fmt.Sprintf("Error accessing registry %v", err))
		return
	}
	return
}

func getRegistryDeleteKey()(key registry.Key, err error) {
	key, _, err = registry.CreateKey(registry.CURRENT_USER, REGISTRY_DELETE_KEY_PATH, registry.SET_VALUE)
	if err != nil {
		logMessage(fmt.Sprintf("Error accessing registry %v", err))
		return
	}
	return
}
