package sbom

import (
	"github.com/CycloneDX/cyclonedx-go"
	cdx "github.com/CycloneDX/cyclonedx-go"
)

func EdgeTypeFromSPDX(spdxName string) Edge_Type {
	switch spdxName {
	case "AMENDS":
		return Edge_amends
	case "ANCESTOR_OF":
		return Edge_ancestor
	case "BUILD_DEPENDENCY_OF":
		return Edge_buildDependency
	case "BUILD_TOOL_OF":
		return Edge_buildTool
		// case "CONTAINED_BY":
	case "CONTAINS":
		return Edge_contains
	case "COPY_OF":
		return Edge_copy
	case "DATA_FILE_OF":
		return Edge_dataFile
	case "DEPENDENCY_MANIFEST_OF":
		return Edge_dependencyManifest
		// case "DEPENDENCY_OF":
	case "DEPENDS_ON":
		return Edge_dependsOn
	case "DESCENDANT_OF":
		return Edge_descendant
		// case "DESCRIBED_BY":
	case "DESCRIBES":
		return Edge_describes
	case "DEV_DEPENDENCY_OF":
		return Edge_devDependency
	case "DEV_TOOL_OF":
		return Edge_devTool
	case "DISTRIBUTION_ARTIFACT":
		return Edge_distributionArtifact
	case "DOCUMENTATION_OF":
		return Edge_documentation
	case "DYNAMIC_LINK":
		return Edge_dynamicLink
	case "EXAMPLE_OF":
		return Edge_example
	case "EXPANDED_FROM_ARCHIVE":
		return Edge_expandedFromArchive
	case "FILE_ADDED":
		return Edge_fileAdded
	case "FILE_DELETED":
		return Edge_fileDeleted
	case "FILE_MODIFIED":
		return Edge_fileModified
		// case "GENERATED_FROM":
	case "GENERATES":
		return Edge_generates
	case "METAFILE_OF":
		return Edge_metafile
	case "OPTIONAL_COMPONENT_OF":
		return Edge_optionalComponent
	case "OPTIONAL_DEPENDENCY_OF":
		return Edge_optionalDependency
	case "OTHER":
		return Edge_other
	case "PACKAGE_OF":
		return Edge_packages
	// case "PATCH_APPLIED":
	case "PATCH_FOR":
		return Edge_patch
	// case "PREREQUISITE_FOR":
	case "HAS_PREREQUISITE":
		return Edge_prerequisite
	case "PROVIDED_DEPENDENCY_OF":
		return Edge_providedDependency
	case "REQUIREMENT_DESCRIPTION_FOR":
		return Edge_requirementFor
	case "RUNTIME_DEPENDENCY_OF":
		return Edge_runtimeDependency
	case "SPECIFICATION_FOR":
		return Edge_specificationFor
	case "STATIC_LINK":
		return Edge_staticLink
	case "TEST_OF":
		return Edge_test
	case "TEST_CASE_OF":
		return Edge_testCase
	case "TEST_DEPENDENCY_OF":
		return Edge_testDependency
	case "TEST_TOOL_OF":
		return Edge_testTool
	case "VARIANT_OF":
		return Edge_variant
	default:
		return Edge_UNKNOWN
	}
}

func HashAlgorithmFromCDX(cdxAlgorithm cyclonedx.HashAlgorithm) HashAlgorithm {
	switch cdxAlgorithm {
	case cdx.HashAlgoMD5:
		return HashAlgorithm_MD5
	case cdx.HashAlgoSHA1:
		return HashAlgorithm_SHA1
	case cdx.HashAlgoSHA256:
		return HashAlgorithm_SHA256
	case cdx.HashAlgoSHA384:
		return HashAlgorithm_SHA384
	case cdx.HashAlgoSHA512:
		return HashAlgorithm_SHA512
	case cdx.HashAlgoSHA3_256:
		return HashAlgorithm_SHA3_256
	case cdx.HashAlgoSHA3_384:
		return HashAlgorithm_SHA3_384
	case cdx.HashAlgoSHA3_512:
		return HashAlgorithm_SHA3_512
	case cdx.HashAlgoBlake2b_256:
		return HashAlgorithm_BLAKE2B_256
	case cdx.HashAlgoBlake2b_384:
		return HashAlgorithm_BLAKE2B_384
	case cdx.HashAlgoBlake2b_512:
		return HashAlgorithm_BLAKE2B_512
	case cdx.HashAlgoBlake3:
		return HashAlgorithm_BLAKE3
	default:
		return HashAlgorithm_UNKNOWN
	}
}
