package checks

var fsCheckList = []check{
	getKernelModuleCheck(
		"1.1.1.1",
		"cramfs",
		"Ensure mounting of cramfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.2",
		"freevxfs",
		"Ensure mounting of freevxfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.3",
		"jffs2",
		"Ensure mounting of jffs2 filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.4",
		"hfs",
		"Ensure mounting of hfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.5",
		"hfsplus",
		"Ensure mounting of hfsplus filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.6",
		"squashfs",
		"Ensure mounting of squashfs filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.7",
		"udf",
		"Ensure mounting of udf filesystems is disabled",
		"filesystem",
		true,
	),
	getKernelModuleCheck(
		"1.1.1.8",
		"vfat",
		"Ensure mounting of FAT filesystems is disabled",
		"filesystem",
		true,
	),
}
