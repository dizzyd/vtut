package main

import (
	"github.com/vulkan-go/demos/vulkaninfo"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

var appInfo = &vk.ApplicationInfo{
	SType:              vk.StructureTypeApplicationInfo,
	ApiVersion:         vk.MakeVersion(1, 0, 0),
	ApplicationVersion: vk.MakeVersion(1, 0, 0),
	PApplicationName:   "VulkanInfo\x00",
	PEngineName:        "vulkango.com\x00",
}

func main() {
	orPanic(glfw.Init())
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())
	orPanic(vk.Init())

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(640, 480, "Vulkan Info", nil, nil)
	orPanic(err)
	defer window.Destroy()

	vkDevice, err := vulkaninfo.NewVulkanDevice(appInfo, window.GLFWWindow())
	orPanic(err)
	vulkaninfo.PrintInfo(vkDevice)
	vkDevice.Destroy()
}
