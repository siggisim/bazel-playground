load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

## Go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "86ae934bd4c43b99893fc64be9d9fc684b81461581df7ea8fc291c816f5ee8c5",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.3/rules_go-0.18.3.tar.gz"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

## Gazelle
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

## Docker
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")

_go_image_repos()

######################################################################################################
### go.mod
######################################################################################################

go_repository(
    name = "co_honnef_go_tools",
    commit = "c2f93a96b099",
    importpath = "honnef.co/go/tools",
)

go_repository(
    name = "com_github_armon_go_metrics",
    commit = "f0300d1749da",
    importpath = "github.com/armon/go-metrics",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    tag = "v1.19.11",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    tag = "v0.3.1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    tag = "v0.3.4",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    tag = "v3.2.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    tag = "v1.2.1",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b",
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    tag = "v0.0.1",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    importpath = "github.com/hashicorp/go-immutable-radix",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    tag = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    tag = "v0.5.0",
)

go_repository(
    name = "com_github_hashicorp_raft",
    importpath = "github.com/hashicorp/raft",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    commit = "c2b33e8439af",
    importpath = "github.com/jmespath/go-jmespath",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    importpath = "github.com/kelseyhightower/envconfig",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_nats_io_gnatsd",
    importpath = "github.com/nats-io/gnatsd",
    tag = "v1.4.1",
)

go_repository(
    name = "com_github_nats_io_go_nats",
    importpath = "github.com/nats-io/go-nats",
    tag = "v1.7.0",
)

go_repository(
    name = "com_github_nats_io_go_nats_streaming",
    importpath = "github.com/nats-io/go-nats-streaming",
    tag = "v0.4.0",
)

go_repository(
    name = "com_github_nats_io_nats_streaming_server",
    importpath = "github.com/nats-io/nats-streaming-server",
    tag = "v0.12.0",
)

go_repository(
    name = "com_github_nats_io_nkeys",
    importpath = "github.com/nats-io/nkeys",
    tag = "v0.0.2",
)

go_repository(
    name = "com_github_nats_io_nuid",
    importpath = "github.com/nats-io/nuid",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_onrik_logrus",
    importpath = "github.com/onrik/logrus",
    tag = "v0.2.2",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    tag = "v1.0.2",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    commit = "57f6aae5913c",
    importpath = "github.com/pascaldekloe/goe",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    tag = "v0.8.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    tag = "v0.9.2",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "5c3871d89910",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "67670fe90761",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "1dc9a6cbc91a",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.3.0",
)

go_repository(
    name = "com_github_utilitywarehouse_go_operational",
    commit = "b3669f314552",
    importpath = "github.com/utilitywarehouse/go-operational",
)

go_repository(
    name = "com_github_utilitywarehouse_go_ops_health_checks",
    commit = "1d2989eb22ad",
    importpath = "github.com/utilitywarehouse/go-ops-health-checks",
)

go_repository(
    name = "com_github_utilitywarehouse_go_pubsub",
    commit = "34bed0f69328",
    importpath = "github.com/utilitywarehouse/go-pubsub",
)

go_repository(
    name = "com_github_utilitywarehouse_partner_pkg_event",
    commit = "81f3c69e9b83",
    importpath = "github.com/utilitywarehouse/partner-pkg/event",
)

go_repository(
    name = "com_github_utilitywarehouse_partner_pkg_log",
    commit = "1ec4478489fd",
    importpath = "github.com/utilitywarehouse/partner-pkg/log",
)

go_repository(
    name = "com_github_utilitywarehouse_partner_pkg_meta",
    commit = "256fb5af71e8",
    importpath = "github.com/utilitywarehouse/partner-pkg/meta",
)

go_repository(
    name = "com_github_utilitywarehouse_partner_pkg_operational",
    commit = "81f3c69e9b83",
    importpath = "github.com/utilitywarehouse/partner-pkg/operational",
)

go_repository(
    name = "com_github_utilitywarehouse_partner_pkg_service",
    commit = "daaf4a5ea56d",
    importpath = "github.com/utilitywarehouse/partner-pkg/service",
)

go_repository(
    name = "com_github_utilitywarehouse_partner_pkg_ticker",
    commit = "6a9cd7133615",
    importpath = "github.com/utilitywarehouse/partner-pkg/ticker",
)

go_repository(
    name = "com_github_uw_labs_freezer",
    commit = "f22cdee36346",
    importpath = "github.com/uw-labs/freezer",
)

go_repository(
    name = "com_github_uw_labs_proximo",
    commit = "db55ed8fbbc7",
    importpath = "github.com/uw-labs/proximo",
)

go_repository(
    name = "com_github_uw_labs_straw",
    commit = "dd4694254d06",
    importpath = "github.com/uw-labs/straw",
)

go_repository(
    name = "com_github_uw_labs_substrate",
    commit = "b796fa07c8c6",
    importpath = "github.com/uw-labs/substrate",
)

go_repository(
    name = "com_github_uw_labs_sync",
    commit = "1bb306bf6e71",
    importpath = "github.com/uw-labs/sync",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    tag = "v0.34.0",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    tag = "v1.3.2",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    tag = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "fc2db5cae922",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    tag = "v1.19.0",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "ff983b9c42bc",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_exp",
    commit = "509febef88a4",
    importpath = "golang.org/x/exp",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "c67002cb31c3",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "org_golang_x_net",
    commit = "c95aed5357e7",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_oauth2",
    commit = "5dab4167f31c",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "37e7f081c4d4",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "7fbe1cd0fcc2",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "bf090417da8b",
    importpath = "golang.org/x/tools",
)
