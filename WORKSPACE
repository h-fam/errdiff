load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")


#http_archive(
#    name = "io_bazel_rules_go",
#    urls = [
#        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
#        "https://github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
#    ],
#    sha256 = "f04d2373bcaf8aa09bccb08a98a57e721306c8f6043a2a0ee610fd6853dcde3d",
#)

git_repository(
    name = "io_bazel_rules_go",
    commit = "fabf03c1cd31bcf15fb945d932cef322b242be3a",
    remote = "https://github.com/bazelbuild/rules_go",
    shallow_since = "1561303606 -0400",
)

#http_archive(
#    name = "bazel_gazelle",
#    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
#    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
#)

git_repository(
    name = "bazel_gazelle",
    shallow_since = "1561306268 -0400",
    commit = "72ba271916ca02aaaff72949dc0c0a63ab37d395",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
)

# proto_library, cc_proto_library, and java_proto_library rules implicitly
# depend on @com_google_protobuf for protoc and proto runtimes.
# This statement defines the @com_google_protobuf repo.
http_archive(
    name = "com_google_protobuf",
    sha256 = "f976a4cd3f1699b6d20c1e944ca1de6754777918320c719742e1674fcf247b7e",
    strip_prefix = "protobuf-3.7.1",
    urls = ["https://github.com/google/protobuf/archive/v3.7.1.zip"],
)

# bazel-skylb 0.8.0 released 2019.03.20 (https://github.com/bazelbuild/bazel-skylib/releases/tag/0.8.0)
skylib_version = "0.8.0"

http_archive(
    name = "bazel_skylib",
    sha256 = "2ef429f5d7ce7111263289644d233707dba35e39696377ebab8b0bc701f7818e",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/{}/bazel-skylib.{}.tar.gz".format(skylib_version, skylib_version),
)

# Add Docker toolchains
http_archive(
    name = "bazel_toolchains",
    sha256 = "56e75f7c9bb074f35b71a9950917fbd036bd1433f9f5be7c04bace0e68eb804a",
    strip_prefix = "bazel-toolchains-9bd2748ec99d72bec41c88eecc3b7bd19d91a0c7",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-toolchains/archive/9bd2748ec99d72bec41c88eecc3b7bd19d91a0c7.tar.gz",
        "https://github.com/bazelbuild/bazel-toolchains/archive/9bd2748ec99d72bec41c88eecc3b7bd19d91a0c7.tar.gz",
    ],
)

# proto_library, cc_proto_library, and java_proto_library rules implicitly
# depend on @com_google_protobuf for protoc and proto runtimes.
# This statement defines the @com_google_protobuf repo.
http_archive(
    name = "com_google_protobuf",
    sha256 = "f976a4cd3f1699b6d20c1e944ca1de6754777918320c719742e1674fcf247b7e",
    strip_prefix = "protobuf-3.7.1",
    urls = ["https://github.com/google/protobuf/archive/v3.7.1.zip"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies(go_sdk = "go_sdk")

# Download using "go mod download"
go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:j6XxA85m/6txkUCHvzlV5f+HBNl/1r5cZ2A/3IEFOO8=",
    version = "v1.21.1",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:4QSRKanuywn15aTZvI/mIDEgPQpswuFndXpOj3rKEco=",
    version = "v0.0.0-20190522155817-f3200d17e092",
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_download_sdk",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_download_sdk(
    name = "go_sdk",
    sdks = {
        "linux_amd64": ("go1.12.6.linux-amd64.tar.gz", "66d83bfb5a9ede000e33c6579a91a29e6b101829ad41fffb5c5bb6c900e109d9"),
    },
)

go_rules_dependencies()

go_register_toolchains()

# Download the rules_docker repository at release v0.7.0
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

# This enables Google Remote Build Execution toolchains
# https://cloud.google.com/remote-build-execution/docs/set-up/remote-environment#add_a_toolchain
load("@bazel_toolchains//rules:rbe_repo.bzl", "rbe_autoconfig")

rbe_autoconfig(name = "rbe_default")
