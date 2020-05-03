package checks

var fsCheckList = []check{
	getKernelModuleCheck(
		"cramfs",
		"Ensure mounting of cramfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"freevxfs",
		"Ensure mounting of freevxfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"jffs2",
		"Ensure mounting of jffs2 filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"hfs",
		"Ensure mounting of hfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"hfsplus",
		"Ensure mounting of hfsplus filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"squashfs",
		"Ensure mounting of squashfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"udf",
		"Ensure mounting of udf filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"vfat",
		"Ensure mounting of FAT filesystems is disabled",
		"filesystem",
		true,
	),
}
