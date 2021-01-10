package(go_import_path = "github.com/tiagovtristao/plz")

go_library(
    name = "importpath",
    srcs = ["importpath.go"],
)

go_test(
    name = "importpath_test",
    srcs = ["importpath_test.go"],
    external = True,
    deps = [
        ":importpath",
        "//third_party/go:testify",
    ],
)
