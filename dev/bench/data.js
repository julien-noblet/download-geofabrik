window.BENCHMARK_DATA = {
  "lastUpdate": 1790291891596,
  "repoUrl": "https://github.com/julien-noblet/download-geofabrik",
  "entries": {
    "download-geofabrik Go Benchmarks": [
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "julien.noblet+github@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "distinct": true,
          "id": "7e8168e2c9a81d782987bf2a24f49a1c4d089ec0",
          "message": "chore(deps): bump docker/login-action from 3 to 4 in the actions group\n\nBumps the actions group with 1 update: [docker/login-action](https://github.com/docker/login-action).\n\n\nUpdates `docker/login-action` from 3 to 4\n- [Release notes](https://github.com/docker/login-action/releases)\n- [Commits](https://github.com/docker/login-action/compare/v3...v4)\n\n---\nupdated-dependencies:\n- dependency-name: docker/login-action\n  dependency-version: '4'\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: actions\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>",
          "timestamp": "2026-09-16T23:34:23+02:00",
          "tree_id": "1902708653c52816e1716f3f9410ccbe11f9c226",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/7e8168e2c9a81d782987bf2a24f49a1c4d089ec0"
        },
        "date": 1789594651884,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 17474,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "58536 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 17474,
            "unit": "ns/op",
            "extra": "58536 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "58536 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "58536 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13471,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "88573 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13471,
            "unit": "ns/op",
            "extra": "88573 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "88573 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "88573 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2056,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "567376 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2056,
            "unit": "ns/op",
            "extra": "567376 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "567376 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "567376 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2052,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "580245 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2052,
            "unit": "ns/op",
            "extra": "580245 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "580245 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "580245 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 226,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5420988 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 226,
            "unit": "ns/op",
            "extra": "5420988 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5420988 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5420988 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 223.6,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5502568 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 223.6,
            "unit": "ns/op",
            "extra": "5502568 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5502568 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5502568 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 361134,
            "unit": "ns/op\t 181.47 MB/s\t    7790 B/op\t      93 allocs/op",
            "extra": "3193 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 361134,
            "unit": "ns/op",
            "extra": "3193 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 181.47,
            "unit": "MB/s",
            "extra": "3193 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7790,
            "unit": "B/op",
            "extra": "3193 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3193 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 360036,
            "unit": "ns/op\t 182.03 MB/s\t    7531 B/op\t      93 allocs/op",
            "extra": "3259 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 360036,
            "unit": "ns/op",
            "extra": "3259 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 182.03,
            "unit": "MB/s",
            "extra": "3259 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7531,
            "unit": "B/op",
            "extra": "3259 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3259 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36477,
            "unit": "ns/op\t   33394 B/op\t       8 allocs/op",
            "extra": "32736 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36477,
            "unit": "ns/op",
            "extra": "32736 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33394,
            "unit": "B/op",
            "extra": "32736 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32736 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36111,
            "unit": "ns/op\t   33333 B/op\t       8 allocs/op",
            "extra": "33267 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36111,
            "unit": "ns/op",
            "extra": "33267 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33333,
            "unit": "B/op",
            "extra": "33267 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "33267 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10199,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "115431 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10199,
            "unit": "ns/op",
            "extra": "115431 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "115431 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "115431 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10317,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "115414 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10317,
            "unit": "ns/op",
            "extra": "115414 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "115414 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "115414 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 137114,
            "unit": "ns/op\t   17408 B/op\t     194 allocs/op",
            "extra": "8803 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 137114,
            "unit": "ns/op",
            "extra": "8803 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17408,
            "unit": "B/op",
            "extra": "8803 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8803 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 137010,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "8934 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 137010,
            "unit": "ns/op",
            "extra": "8934 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "8934 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8934 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 438158,
            "unit": "ns/op\t   73866 B/op\t     648 allocs/op",
            "extra": "2629 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 438158,
            "unit": "ns/op",
            "extra": "2629 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73866,
            "unit": "B/op",
            "extra": "2629 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2629 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 446302,
            "unit": "ns/op\t   73811 B/op\t     648 allocs/op",
            "extra": "2719 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 446302,
            "unit": "ns/op",
            "extra": "2719 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73811,
            "unit": "B/op",
            "extra": "2719 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2719 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 136946,
            "unit": "ns/op\t   14736 B/op\t     109 allocs/op",
            "extra": "8247 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 136946,
            "unit": "ns/op",
            "extra": "8247 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14736,
            "unit": "B/op",
            "extra": "8247 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8247 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 135448,
            "unit": "ns/op\t   14716 B/op\t     109 allocs/op",
            "extra": "8902 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 135448,
            "unit": "ns/op",
            "extra": "8902 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14716,
            "unit": "B/op",
            "extra": "8902 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8902 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 133005,
            "unit": "ns/op\t   13043 B/op\t     104 allocs/op",
            "extra": "8980 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 133005,
            "unit": "ns/op",
            "extra": "8980 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13043,
            "unit": "B/op",
            "extra": "8980 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8980 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 132932,
            "unit": "ns/op\t   13027 B/op\t     104 allocs/op",
            "extra": "9469 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 132932,
            "unit": "ns/op",
            "extra": "9469 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13027,
            "unit": "B/op",
            "extra": "9469 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "9469 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 547329,
            "unit": "ns/op\t   82321 B/op\t     678 allocs/op",
            "extra": "2148 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 547329,
            "unit": "ns/op",
            "extra": "2148 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82321,
            "unit": "B/op",
            "extra": "2148 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2148 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 547734,
            "unit": "ns/op\t   82264 B/op\t     678 allocs/op",
            "extra": "2170 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 547734,
            "unit": "ns/op",
            "extra": "2170 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82264,
            "unit": "B/op",
            "extra": "2170 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2170 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 141526,
            "unit": "ns/op\t   15560 B/op\t     224 allocs/op",
            "extra": "8404 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 141526,
            "unit": "ns/op",
            "extra": "8404 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15560,
            "unit": "B/op",
            "extra": "8404 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8404 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 140286,
            "unit": "ns/op\t   15546 B/op\t     224 allocs/op",
            "extra": "8920 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 140286,
            "unit": "ns/op",
            "extra": "8920 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15546,
            "unit": "B/op",
            "extra": "8920 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8920 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 267373,
            "unit": "ns/op\t   27399 B/op\t     354 allocs/op",
            "extra": "4323 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 267373,
            "unit": "ns/op",
            "extra": "4323 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27399,
            "unit": "B/op",
            "extra": "4323 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4323 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 267572,
            "unit": "ns/op\t   27380 B/op\t     354 allocs/op",
            "extra": "4507 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 267572,
            "unit": "ns/op",
            "extra": "4507 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27380,
            "unit": "B/op",
            "extra": "4507 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4507 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 384142,
            "unit": "ns/op\t  947884 B/op\t    1256 allocs/op",
            "extra": "2928 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 384142,
            "unit": "ns/op",
            "extra": "2928 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947884,
            "unit": "B/op",
            "extra": "2928 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "2928 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 386763,
            "unit": "ns/op\t  947954 B/op\t    1256 allocs/op",
            "extra": "2708 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 386763,
            "unit": "ns/op",
            "extra": "2708 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947954,
            "unit": "B/op",
            "extra": "2708 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "2708 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 136012,
            "unit": "ns/op\t   15127 B/op\t     208 allocs/op",
            "extra": "8689 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 136012,
            "unit": "ns/op",
            "extra": "8689 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15127,
            "unit": "B/op",
            "extra": "8689 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "8689 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 136699,
            "unit": "ns/op\t   15114 B/op\t     208 allocs/op",
            "extra": "8228 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 136699,
            "unit": "ns/op",
            "extra": "8228 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15114,
            "unit": "B/op",
            "extra": "8228 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "8228 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 131116,
            "unit": "ns/op\t   14518 B/op\t     174 allocs/op",
            "extra": "9456 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 131116,
            "unit": "ns/op",
            "extra": "9456 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14518,
            "unit": "B/op",
            "extra": "9456 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9456 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 130521,
            "unit": "ns/op\t   14506 B/op\t     174 allocs/op",
            "extra": "9688 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 130521,
            "unit": "ns/op",
            "extra": "9688 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14506,
            "unit": "B/op",
            "extra": "9688 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9688 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 540685,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2200 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 540685,
            "unit": "ns/op",
            "extra": "2200 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2200 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2200 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 542510,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 542510,
            "unit": "ns/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2211 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 352879,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3446 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 352879,
            "unit": "ns/op",
            "extra": "3446 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3446 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3446 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 338144,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3457 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 338144,
            "unit": "ns/op",
            "extra": "3457 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3457 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3457 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 952632,
            "unit": "ns/op\t  335370 B/op\t     573 allocs/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 952632,
            "unit": "ns/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335370,
            "unit": "B/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 951320,
            "unit": "ns/op\t  332748 B/op\t     573 allocs/op",
            "extra": "1236 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 951320,
            "unit": "ns/op",
            "extra": "1236 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 332748,
            "unit": "B/op",
            "extra": "1236 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1236 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 30.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38121289 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 30.55,
            "unit": "ns/op",
            "extra": "38121289 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38121289 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38121289 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.77,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44897427 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.77,
            "unit": "ns/op",
            "extra": "44897427 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44897427 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44897427 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 100.8,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11856078 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 100.8,
            "unit": "ns/op",
            "extra": "11856078 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11856078 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11856078 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 103,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11634454 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 103,
            "unit": "ns/op",
            "extra": "11634454 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11634454 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11634454 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 127.9,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9744529 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 127.9,
            "unit": "ns/op",
            "extra": "9744529 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9744529 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9744529 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 128.9,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "7891260 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 128.9,
            "unit": "ns/op",
            "extra": "7891260 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "7891260 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7891260 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 327.6,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3665889 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 327.6,
            "unit": "ns/op",
            "extra": "3665889 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3665889 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3665889 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 328.5,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3634490 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 328.5,
            "unit": "ns/op",
            "extra": "3634490 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3634490 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3634490 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73237,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16346 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73237,
            "unit": "ns/op",
            "extra": "16346 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16346 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16346 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 72445,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16540 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 72445,
            "unit": "ns/op",
            "extra": "16540 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16540 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16540 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8972108,
            "unit": "ns/op\t 2984569 B/op\t   60446 allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8972108,
            "unit": "ns/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984569,
            "unit": "B/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8977262,
            "unit": "ns/op\t 2984570 B/op\t   60446 allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8977262,
            "unit": "ns/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984570,
            "unit": "B/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10067319,
            "unit": "ns/op\t15163500 B/op\t   23990 allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10067319,
            "unit": "ns/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15163500,
            "unit": "B/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23990,
            "unit": "allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10380891,
            "unit": "ns/op\t15164160 B/op\t   23994 allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10380891,
            "unit": "ns/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164160,
            "unit": "B/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23994,
            "unit": "allocs/op",
            "extra": "100 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "julien.noblet+github@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "distinct": false,
          "id": "7e8168e2c9a81d782987bf2a24f49a1c4d089ec0",
          "message": "chore(deps): bump docker/login-action from 3 to 4 in the actions group\n\nBumps the actions group with 1 update: [docker/login-action](https://github.com/docker/login-action).\n\n\nUpdates `docker/login-action` from 3 to 4\n- [Release notes](https://github.com/docker/login-action/releases)\n- [Commits](https://github.com/docker/login-action/compare/v3...v4)\n\n---\nupdated-dependencies:\n- dependency-name: docker/login-action\n  dependency-version: '4'\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: actions\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>",
          "timestamp": "2026-09-16T23:34:23+02:00",
          "tree_id": "1902708653c52816e1716f3f9410ccbe11f9c226",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/7e8168e2c9a81d782987bf2a24f49a1c4d089ec0"
        },
        "date": 1789594797408,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 10582,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "111450 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 10582,
            "unit": "ns/op",
            "extra": "111450 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "111450 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "111450 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 10012,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "120403 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 10012,
            "unit": "ns/op",
            "extra": "120403 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "120403 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "120403 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 689.8,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "1742412 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 689.8,
            "unit": "ns/op",
            "extra": "1742412 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1742412 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1742412 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 691.3,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "1729107 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 691.3,
            "unit": "ns/op",
            "extra": "1729107 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1729107 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1729107 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 205.1,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5861832 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 205.1,
            "unit": "ns/op",
            "extra": "5861832 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5861832 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5861832 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 216.7,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5865079 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 216.7,
            "unit": "ns/op",
            "extra": "5865079 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5865079 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5865079 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 234350,
            "unit": "ns/op\t 279.65 MB/s\t    7685 B/op\t      93 allocs/op",
            "extra": "5221 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 234350,
            "unit": "ns/op",
            "extra": "5221 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 279.65,
            "unit": "MB/s",
            "extra": "5221 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7685,
            "unit": "B/op",
            "extra": "5221 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "5221 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 233890,
            "unit": "ns/op\t 280.20 MB/s\t    7535 B/op\t      93 allocs/op",
            "extra": "4966 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 233890,
            "unit": "ns/op",
            "extra": "4966 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 280.2,
            "unit": "MB/s",
            "extra": "4966 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7535,
            "unit": "B/op",
            "extra": "4966 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "4966 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 27960,
            "unit": "ns/op\t   33349 B/op\t       8 allocs/op",
            "extra": "42804 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 27960,
            "unit": "ns/op",
            "extra": "42804 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33349,
            "unit": "B/op",
            "extra": "42804 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "42804 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 28062,
            "unit": "ns/op\t   33322 B/op\t       8 allocs/op",
            "extra": "42680 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 28062,
            "unit": "ns/op",
            "extra": "42680 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33322,
            "unit": "B/op",
            "extra": "42680 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "42680 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 4190,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "286155 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 4190,
            "unit": "ns/op",
            "extra": "286155 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "286155 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "286155 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 4343,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "285249 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 4343,
            "unit": "ns/op",
            "extra": "285249 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "285249 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "285249 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 70150,
            "unit": "ns/op\t   17400 B/op\t     194 allocs/op",
            "extra": "16762 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 70150,
            "unit": "ns/op",
            "extra": "16762 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17400,
            "unit": "B/op",
            "extra": "16762 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "16762 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 69046,
            "unit": "ns/op\t   17393 B/op\t     194 allocs/op",
            "extra": "17548 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 69046,
            "unit": "ns/op",
            "extra": "17548 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17393,
            "unit": "B/op",
            "extra": "17548 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "17548 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 244439,
            "unit": "ns/op\t   73846 B/op\t     648 allocs/op",
            "extra": "4119 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 244439,
            "unit": "ns/op",
            "extra": "4119 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73846,
            "unit": "B/op",
            "extra": "4119 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "4119 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 244481,
            "unit": "ns/op\t   73813 B/op\t     648 allocs/op",
            "extra": "4484 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 244481,
            "unit": "ns/op",
            "extra": "4484 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73813,
            "unit": "B/op",
            "extra": "4484 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "4484 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 68074,
            "unit": "ns/op\t   14725 B/op\t     109 allocs/op",
            "extra": "17520 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 68074,
            "unit": "ns/op",
            "extra": "17520 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14725,
            "unit": "B/op",
            "extra": "17520 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "17520 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 68512,
            "unit": "ns/op\t   14715 B/op\t     109 allocs/op",
            "extra": "17376 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 68512,
            "unit": "ns/op",
            "extra": "17376 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14715,
            "unit": "B/op",
            "extra": "17376 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "17376 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 69336,
            "unit": "ns/op\t   13034 B/op\t     104 allocs/op",
            "extra": "16956 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 69336,
            "unit": "ns/op",
            "extra": "16956 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13034,
            "unit": "B/op",
            "extra": "16956 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "16956 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 67071,
            "unit": "ns/op\t   13024 B/op\t     104 allocs/op",
            "extra": "18002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 67071,
            "unit": "ns/op",
            "extra": "18002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13024,
            "unit": "B/op",
            "extra": "18002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "18002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 297694,
            "unit": "ns/op\t   82286 B/op\t     678 allocs/op",
            "extra": "3656 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 297694,
            "unit": "ns/op",
            "extra": "3656 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82286,
            "unit": "B/op",
            "extra": "3656 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "3656 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 297002,
            "unit": "ns/op\t   82254 B/op\t     678 allocs/op",
            "extra": "3980 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 297002,
            "unit": "ns/op",
            "extra": "3980 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82254,
            "unit": "B/op",
            "extra": "3980 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "3980 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 70990,
            "unit": "ns/op\t   15554 B/op\t     224 allocs/op",
            "extra": "16768 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 70990,
            "unit": "ns/op",
            "extra": "16768 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15554,
            "unit": "B/op",
            "extra": "16768 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "16768 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 68323,
            "unit": "ns/op\t   15544 B/op\t     224 allocs/op",
            "extra": "17913 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 68323,
            "unit": "ns/op",
            "extra": "17913 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15544,
            "unit": "B/op",
            "extra": "17913 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "17913 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 137250,
            "unit": "ns/op\t   27382 B/op\t     354 allocs/op",
            "extra": "8498 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 137250,
            "unit": "ns/op",
            "extra": "8498 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27382,
            "unit": "B/op",
            "extra": "8498 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "8498 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 139964,
            "unit": "ns/op\t   27368 B/op\t     354 allocs/op",
            "extra": "9609 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 139964,
            "unit": "ns/op",
            "extra": "9609 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27368,
            "unit": "B/op",
            "extra": "9609 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "9609 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 324074,
            "unit": "ns/op\t  949358 B/op\t    1256 allocs/op",
            "extra": "3526 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 324074,
            "unit": "ns/op",
            "extra": "3526 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 949358,
            "unit": "B/op",
            "extra": "3526 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3526 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 322749,
            "unit": "ns/op\t  948950 B/op\t    1256 allocs/op",
            "extra": "3514 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 322749,
            "unit": "ns/op",
            "extra": "3514 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 948950,
            "unit": "B/op",
            "extra": "3514 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3514 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 71286,
            "unit": "ns/op\t   15120 B/op\t     208 allocs/op",
            "extra": "16824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 71286,
            "unit": "ns/op",
            "extra": "16824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15120,
            "unit": "B/op",
            "extra": "16824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "16824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 69240,
            "unit": "ns/op\t   15112 B/op\t     208 allocs/op",
            "extra": "17587 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 69240,
            "unit": "ns/op",
            "extra": "17587 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "17587 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "17587 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 67148,
            "unit": "ns/op\t   14512 B/op\t     174 allocs/op",
            "extra": "17908 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 67148,
            "unit": "ns/op",
            "extra": "17908 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14512,
            "unit": "B/op",
            "extra": "17908 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "17908 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 64674,
            "unit": "ns/op\t   14505 B/op\t     174 allocs/op",
            "extra": "18000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 64674,
            "unit": "ns/op",
            "extra": "18000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14505,
            "unit": "B/op",
            "extra": "18000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "18000 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 438727,
            "unit": "ns/op\t  290019 B/op\t    3969 allocs/op",
            "extra": "2647 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 438727,
            "unit": "ns/op",
            "extra": "2647 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290019,
            "unit": "B/op",
            "extra": "2647 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2647 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 437202,
            "unit": "ns/op\t  290019 B/op\t    3969 allocs/op",
            "extra": "2736 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 437202,
            "unit": "ns/op",
            "extra": "2736 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290019,
            "unit": "B/op",
            "extra": "2736 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2736 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 276097,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "4293 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 276097,
            "unit": "ns/op",
            "extra": "4293 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "4293 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "4293 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 276950,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "4174 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 276950,
            "unit": "ns/op",
            "extra": "4174 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "4174 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "4174 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 755172,
            "unit": "ns/op\t  332587 B/op\t     573 allocs/op",
            "extra": "1470 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 755172,
            "unit": "ns/op",
            "extra": "1470 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 332587,
            "unit": "B/op",
            "extra": "1470 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1470 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 764405,
            "unit": "ns/op\t  339801 B/op\t     573 allocs/op",
            "extra": "1563 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 764405,
            "unit": "ns/op",
            "extra": "1563 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 339801,
            "unit": "B/op",
            "extra": "1563 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1563 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 24.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48362870 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 24.62,
            "unit": "ns/op",
            "extra": "48362870 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48362870 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48362870 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 24.66,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48656694 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 24.66,
            "unit": "ns/op",
            "extra": "48656694 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48656694 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48656694 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 79.15,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "15326893 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 79.15,
            "unit": "ns/op",
            "extra": "15326893 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "15326893 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "15326893 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 75.15,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "15876517 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 75.15,
            "unit": "ns/op",
            "extra": "15876517 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "15876517 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "15876517 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 99.12,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "12368984 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 99.12,
            "unit": "ns/op",
            "extra": "12368984 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "12368984 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "12368984 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 98.29,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "11996586 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 98.29,
            "unit": "ns/op",
            "extra": "11996586 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "11996586 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11996586 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 234.3,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "5136816 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 234.3,
            "unit": "ns/op",
            "extra": "5136816 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5136816 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "5136816 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 233.9,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "5134940 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 233.9,
            "unit": "ns/op",
            "extra": "5134940 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "5134940 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "5134940 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 68209,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "17580 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 68209,
            "unit": "ns/op",
            "extra": "17580 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "17580 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17580 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 67812,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "17672 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 67812,
            "unit": "ns/op",
            "extra": "17672 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "17672 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17672 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 6520759,
            "unit": "ns/op\t 2984581 B/op\t   60446 allocs/op",
            "extra": "183 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 6520759,
            "unit": "ns/op",
            "extra": "183 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984581,
            "unit": "B/op",
            "extra": "183 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "183 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 6504793,
            "unit": "ns/op\t 2984566 B/op\t   60446 allocs/op",
            "extra": "184 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 6504793,
            "unit": "ns/op",
            "extra": "184 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984566,
            "unit": "B/op",
            "extra": "184 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "184 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8555753,
            "unit": "ns/op\t15165051 B/op\t   23999 allocs/op",
            "extra": "140 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8555753,
            "unit": "ns/op",
            "extra": "140 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15165051,
            "unit": "B/op",
            "extra": "140 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23999,
            "unit": "allocs/op",
            "extra": "140 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8049878,
            "unit": "ns/op\t15164160 B/op\t   23994 allocs/op",
            "extra": "147 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8049878,
            "unit": "ns/op",
            "extra": "147 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164160,
            "unit": "B/op",
            "extra": "147 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23994,
            "unit": "allocs/op",
            "extra": "147 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julien.noblet@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "committer": {
            "email": "julien.noblet@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "distinct": false,
          "id": "19f636507206e6b29b3033dc9fe9d87fa4c02df6",
          "message": "ci(goreleaser): set up QEMU and Docker Buildx for container attestation",
          "timestamp": "2026-09-16T23:55:07+02:00",
          "tree_id": "2d7e314a7e9aa9763e73a0ff2c8c4fbb6c4a2f72",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/19f636507206e6b29b3033dc9fe9d87fa4c02df6"
        },
        "date": 1789596029439,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13633,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "83877 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13633,
            "unit": "ns/op",
            "extra": "83877 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "83877 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "83877 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 12578,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "94200 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 12578,
            "unit": "ns/op",
            "extra": "94200 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "94200 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "94200 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2424,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "512928 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2424,
            "unit": "ns/op",
            "extra": "512928 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "512928 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "512928 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2387,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "473454 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2387,
            "unit": "ns/op",
            "extra": "473454 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "473454 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "473454 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 224.9,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5345725 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 224.9,
            "unit": "ns/op",
            "extra": "5345725 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5345725 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5345725 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 224.5,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5339073 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 224.5,
            "unit": "ns/op",
            "extra": "5339073 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5339073 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5339073 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 368585,
            "unit": "ns/op\t 177.80 MB/s\t    7745 B/op\t      93 allocs/op",
            "extra": "3271 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 368585,
            "unit": "ns/op",
            "extra": "3271 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 177.8,
            "unit": "MB/s",
            "extra": "3271 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7745,
            "unit": "B/op",
            "extra": "3271 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3271 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 363350,
            "unit": "ns/op\t 180.37 MB/s\t    7546 B/op\t      93 allocs/op",
            "extra": "3291 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 363350,
            "unit": "ns/op",
            "extra": "3291 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 180.37,
            "unit": "MB/s",
            "extra": "3291 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7546,
            "unit": "B/op",
            "extra": "3291 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3291 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 40784,
            "unit": "ns/op\t   33457 B/op\t       8 allocs/op",
            "extra": "29298 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 40784,
            "unit": "ns/op",
            "extra": "29298 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33457,
            "unit": "B/op",
            "extra": "29298 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "29298 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 40497,
            "unit": "ns/op\t   33370 B/op\t       8 allocs/op",
            "extra": "29647 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 40497,
            "unit": "ns/op",
            "extra": "29647 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33370,
            "unit": "B/op",
            "extra": "29647 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "29647 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 13101,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "93220 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 13101,
            "unit": "ns/op",
            "extra": "93220 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "93220 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "93220 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 13350,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "94281 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 13350,
            "unit": "ns/op",
            "extra": "94281 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "94281 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "94281 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 105580,
            "unit": "ns/op\t   17408 B/op\t     194 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 105580,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17408,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 107704,
            "unit": "ns/op\t   17395 B/op\t     194 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 107704,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17395,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 355033,
            "unit": "ns/op\t   73857 B/op\t     648 allocs/op",
            "extra": "3015 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 355033,
            "unit": "ns/op",
            "extra": "3015 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73857,
            "unit": "B/op",
            "extra": "3015 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3015 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 379895,
            "unit": "ns/op\t   73808 B/op\t     648 allocs/op",
            "extra": "3000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 379895,
            "unit": "ns/op",
            "extra": "3000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73808,
            "unit": "B/op",
            "extra": "3000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 111205,
            "unit": "ns/op\t   14733 B/op\t     109 allocs/op",
            "extra": "9962 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 111205,
            "unit": "ns/op",
            "extra": "9962 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14733,
            "unit": "B/op",
            "extra": "9962 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9962 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 113382,
            "unit": "ns/op\t   14714 B/op\t     109 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 113382,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14714,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 111693,
            "unit": "ns/op\t   13042 B/op\t     104 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 111693,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13042,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 109401,
            "unit": "ns/op\t   13027 B/op\t     104 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 109401,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13027,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 460311,
            "unit": "ns/op\t   82310 B/op\t     678 allocs/op",
            "extra": "2462 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 460311,
            "unit": "ns/op",
            "extra": "2462 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82310,
            "unit": "B/op",
            "extra": "2462 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2462 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 454012,
            "unit": "ns/op\t   82260 B/op\t     678 allocs/op",
            "extra": "2624 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 454012,
            "unit": "ns/op",
            "extra": "2624 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82260,
            "unit": "B/op",
            "extra": "2624 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2624 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 108327,
            "unit": "ns/op\t   15558 B/op\t     224 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 108327,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15558,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 109021,
            "unit": "ns/op\t   15545 B/op\t     224 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 109021,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15545,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 212778,
            "unit": "ns/op\t   27389 B/op\t     354 allocs/op",
            "extra": "5713 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 212778,
            "unit": "ns/op",
            "extra": "5713 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27389,
            "unit": "B/op",
            "extra": "5713 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "5713 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 210143,
            "unit": "ns/op\t   27369 B/op\t     354 allocs/op",
            "extra": "5774 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 210143,
            "unit": "ns/op",
            "extra": "5774 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27369,
            "unit": "B/op",
            "extra": "5774 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "5774 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 353160,
            "unit": "ns/op\t  944926 B/op\t    1255 allocs/op",
            "extra": "3394 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 353160,
            "unit": "ns/op",
            "extra": "3394 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 944926,
            "unit": "B/op",
            "extra": "3394 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3394 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 362933,
            "unit": "ns/op\t  944416 B/op\t    1255 allocs/op",
            "extra": "3320 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 362933,
            "unit": "ns/op",
            "extra": "3320 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 944416,
            "unit": "B/op",
            "extra": "3320 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3320 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 107542,
            "unit": "ns/op\t   15127 B/op\t     208 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 107542,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15127,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 108023,
            "unit": "ns/op\t   15113 B/op\t     208 allocs/op",
            "extra": "9685 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 108023,
            "unit": "ns/op",
            "extra": "9685 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15113,
            "unit": "B/op",
            "extra": "9685 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9685 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 102790,
            "unit": "ns/op\t   14518 B/op\t     174 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 102790,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14518,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 100011,
            "unit": "ns/op\t   14505 B/op\t     174 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 100011,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14505,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 519103,
            "unit": "ns/op\t  290019 B/op\t    3969 allocs/op",
            "extra": "2282 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 519103,
            "unit": "ns/op",
            "extra": "2282 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290019,
            "unit": "B/op",
            "extra": "2282 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2282 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 514230,
            "unit": "ns/op\t  290018 B/op\t    3969 allocs/op",
            "extra": "2348 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 514230,
            "unit": "ns/op",
            "extra": "2348 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290018,
            "unit": "B/op",
            "extra": "2348 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2348 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 321836,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3567 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 321836,
            "unit": "ns/op",
            "extra": "3567 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3567 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3567 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 325639,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3546 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 325639,
            "unit": "ns/op",
            "extra": "3546 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3546 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3546 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 909917,
            "unit": "ns/op\t  331418 B/op\t     573 allocs/op",
            "extra": "1299 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 909917,
            "unit": "ns/op",
            "extra": "1299 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 331418,
            "unit": "B/op",
            "extra": "1299 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1299 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 928820,
            "unit": "ns/op\t  333656 B/op\t     573 allocs/op",
            "extra": "1279 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 928820,
            "unit": "ns/op",
            "extra": "1279 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 333656,
            "unit": "B/op",
            "extra": "1279 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1279 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44887015 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.62,
            "unit": "ns/op",
            "extra": "44887015 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44887015 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44887015 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 27.45,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "43615158 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 27.45,
            "unit": "ns/op",
            "extra": "43615158 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "43615158 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "43615158 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 112.4,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "10666927 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 112.4,
            "unit": "ns/op",
            "extra": "10666927 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "10666927 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "10666927 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 112.6,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "10575102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 112.6,
            "unit": "ns/op",
            "extra": "10575102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "10575102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "10575102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 119,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "10184307 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 119,
            "unit": "ns/op",
            "extra": "10184307 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "10184307 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10184307 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 117.6,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9824476 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 117.6,
            "unit": "ns/op",
            "extra": "9824476 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9824476 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9824476 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 327.6,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3666235 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 327.6,
            "unit": "ns/op",
            "extra": "3666235 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3666235 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3666235 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 327.3,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3677742 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 327.3,
            "unit": "ns/op",
            "extra": "3677742 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3677742 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3677742 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73811,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16192 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73811,
            "unit": "ns/op",
            "extra": "16192 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16192 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16192 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73879,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16218 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73879,
            "unit": "ns/op",
            "extra": "16218 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16218 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16218 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 7834250,
            "unit": "ns/op\t 2984570 B/op\t   60446 allocs/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 7834250,
            "unit": "ns/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984570,
            "unit": "B/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 7777039,
            "unit": "ns/op\t 2984570 B/op\t   60446 allocs/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 7777039,
            "unit": "ns/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984570,
            "unit": "B/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10354316,
            "unit": "ns/op\t15164037 B/op\t   23993 allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10354316,
            "unit": "ns/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164037,
            "unit": "B/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23993,
            "unit": "allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10318527,
            "unit": "ns/op\t15164750 B/op\t   23997 allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10318527,
            "unit": "ns/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164750,
            "unit": "B/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23997,
            "unit": "allocs/op",
            "extra": "100 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ccf6440c5091ca6235ec094e0c7cb0775e95564a",
          "message": "chore(deps): update docker/setup-qemu-action action to v4",
          "timestamp": "2026-09-17T03:35:40Z",
          "tree_id": "89072cdb2c9e844ca1e148755fa6cf6bce59e7e9",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/ccf6440c5091ca6235ec094e0c7cb0775e95564a"
        },
        "date": 1789616325147,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 14241,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "79605 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 14241,
            "unit": "ns/op",
            "extra": "79605 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "79605 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "79605 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13100,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "92118 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13100,
            "unit": "ns/op",
            "extra": "92118 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "92118 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "92118 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 938.2,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "1277941 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 938.2,
            "unit": "ns/op",
            "extra": "1277941 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1277941 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1277941 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 941.4,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "1273090 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 941.4,
            "unit": "ns/op",
            "extra": "1273090 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1273090 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1273090 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 250.3,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "4645084 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 250.3,
            "unit": "ns/op",
            "extra": "4645084 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "4645084 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "4645084 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 237.2,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5086639 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 237.2,
            "unit": "ns/op",
            "extra": "5086639 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5086639 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5086639 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 291728,
            "unit": "ns/op\t 224.65 MB/s\t    7712 B/op\t      93 allocs/op",
            "extra": "4005 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 291728,
            "unit": "ns/op",
            "extra": "4005 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 224.65,
            "unit": "MB/s",
            "extra": "4005 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7712,
            "unit": "B/op",
            "extra": "4005 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "4005 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 295549,
            "unit": "ns/op\t 221.74 MB/s\t    7434 B/op\t      93 allocs/op",
            "extra": "4086 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 295549,
            "unit": "ns/op",
            "extra": "4086 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 221.74,
            "unit": "MB/s",
            "extra": "4086 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7434,
            "unit": "B/op",
            "extra": "4086 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "4086 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 30902,
            "unit": "ns/op\t   33369 B/op\t       8 allocs/op",
            "extra": "38365 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 30902,
            "unit": "ns/op",
            "extra": "38365 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33369,
            "unit": "B/op",
            "extra": "38365 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "38365 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 31344,
            "unit": "ns/op\t   33365 B/op\t       8 allocs/op",
            "extra": "38396 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 31344,
            "unit": "ns/op",
            "extra": "38396 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33365,
            "unit": "B/op",
            "extra": "38396 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "38396 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 5388,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "221865 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 5388,
            "unit": "ns/op",
            "extra": "221865 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "221865 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "221865 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 5389,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "213919 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 5389,
            "unit": "ns/op",
            "extra": "213919 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "213919 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "213919 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 102551,
            "unit": "ns/op\t   17407 B/op\t     194 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 102551,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17407,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 103070,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 103070,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 325249,
            "unit": "ns/op\t   73855 B/op\t     648 allocs/op",
            "extra": "3396 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 325249,
            "unit": "ns/op",
            "extra": "3396 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73855,
            "unit": "B/op",
            "extra": "3396 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3396 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 334381,
            "unit": "ns/op\t   73815 B/op\t     648 allocs/op",
            "extra": "3534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 334381,
            "unit": "ns/op",
            "extra": "3534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73815,
            "unit": "B/op",
            "extra": "3534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 103544,
            "unit": "ns/op\t   14734 B/op\t     109 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 103544,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14734,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 102088,
            "unit": "ns/op\t   14717 B/op\t     109 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 102088,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14717,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 98886,
            "unit": "ns/op\t   13040 B/op\t     104 allocs/op",
            "extra": "12127 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 98886,
            "unit": "ns/op",
            "extra": "12127 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13040,
            "unit": "B/op",
            "extra": "12127 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "12127 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 97056,
            "unit": "ns/op\t   13027 B/op\t     104 allocs/op",
            "extra": "12381 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 97056,
            "unit": "ns/op",
            "extra": "12381 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13027,
            "unit": "B/op",
            "extra": "12381 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "12381 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 409726,
            "unit": "ns/op\t   82300 B/op\t     678 allocs/op",
            "extra": "2893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 409726,
            "unit": "ns/op",
            "extra": "2893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82300,
            "unit": "B/op",
            "extra": "2893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 417545,
            "unit": "ns/op\t   82256 B/op\t     678 allocs/op",
            "extra": "2844 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 417545,
            "unit": "ns/op",
            "extra": "2844 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82256,
            "unit": "B/op",
            "extra": "2844 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2844 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 103721,
            "unit": "ns/op\t   15557 B/op\t     224 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 103721,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15557,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 100813,
            "unit": "ns/op\t   15545 B/op\t     224 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 100813,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15545,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 186890,
            "unit": "ns/op\t   27389 B/op\t     354 allocs/op",
            "extra": "6105 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 186890,
            "unit": "ns/op",
            "extra": "6105 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27389,
            "unit": "B/op",
            "extra": "6105 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "6105 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 197596,
            "unit": "ns/op\t   27371 B/op\t     354 allocs/op",
            "extra": "6020 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 197596,
            "unit": "ns/op",
            "extra": "6020 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27371,
            "unit": "B/op",
            "extra": "6020 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "6020 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 352432,
            "unit": "ns/op\t  947549 B/op\t    1255 allocs/op",
            "extra": "3382 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 352432,
            "unit": "ns/op",
            "extra": "3382 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947549,
            "unit": "B/op",
            "extra": "3382 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3382 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 359399,
            "unit": "ns/op\t  947129 B/op\t    1255 allocs/op",
            "extra": "3453 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 359399,
            "unit": "ns/op",
            "extra": "3453 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947129,
            "unit": "B/op",
            "extra": "3453 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3453 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 100197,
            "unit": "ns/op\t   15125 B/op\t     208 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 100197,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 101002,
            "unit": "ns/op\t   15114 B/op\t     208 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 101002,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15114,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 96276,
            "unit": "ns/op\t   14516 B/op\t     174 allocs/op",
            "extra": "12409 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 96276,
            "unit": "ns/op",
            "extra": "12409 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14516,
            "unit": "B/op",
            "extra": "12409 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "12409 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 96515,
            "unit": "ns/op\t   14505 B/op\t     174 allocs/op",
            "extra": "12537 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 96515,
            "unit": "ns/op",
            "extra": "12537 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14505,
            "unit": "B/op",
            "extra": "12537 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "12537 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 537697,
            "unit": "ns/op\t  290014 B/op\t    3969 allocs/op",
            "extra": "2187 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 537697,
            "unit": "ns/op",
            "extra": "2187 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290014,
            "unit": "B/op",
            "extra": "2187 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2187 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 547484,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2246 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 547484,
            "unit": "ns/op",
            "extra": "2246 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2246 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2246 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 333103,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3570 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 333103,
            "unit": "ns/op",
            "extra": "3570 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3570 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3570 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 333121,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3516 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 333121,
            "unit": "ns/op",
            "extra": "3516 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3516 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3516 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 904225,
            "unit": "ns/op\t  335505 B/op\t     573 allocs/op",
            "extra": "1303 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 904225,
            "unit": "ns/op",
            "extra": "1303 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335505,
            "unit": "B/op",
            "extra": "1303 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1303 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 900973,
            "unit": "ns/op\t  332911 B/op\t     573 allocs/op",
            "extra": "1305 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 900973,
            "unit": "ns/op",
            "extra": "1305 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 332911,
            "unit": "B/op",
            "extra": "1305 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1305 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44653658 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.7,
            "unit": "ns/op",
            "extra": "44653658 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44653658 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44653658 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.88,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44986926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.88,
            "unit": "ns/op",
            "extra": "44986926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44986926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44986926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 94.08,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "12380558 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 94.08,
            "unit": "ns/op",
            "extra": "12380558 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "12380558 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12380558 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 93.49,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "12797102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 93.49,
            "unit": "ns/op",
            "extra": "12797102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "12797102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12797102 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.7,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9322339 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.7,
            "unit": "ns/op",
            "extra": "9322339 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9322339 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9322339 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.8,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9379238 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.8,
            "unit": "ns/op",
            "extra": "9379238 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9379238 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9379238 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 303.4,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3941188 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 303.4,
            "unit": "ns/op",
            "extra": "3941188 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3941188 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3941188 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 307.4,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3859851 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 307.4,
            "unit": "ns/op",
            "extra": "3859851 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3859851 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3859851 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 74234,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16166 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 74234,
            "unit": "ns/op",
            "extra": "16166 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16166 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16166 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 74494,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16082 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 74494,
            "unit": "ns/op",
            "extra": "16082 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16082 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16082 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8365027,
            "unit": "ns/op\t 2984574 B/op\t   60446 allocs/op",
            "extra": "144 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8365027,
            "unit": "ns/op",
            "extra": "144 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984574,
            "unit": "B/op",
            "extra": "144 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "144 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8271608,
            "unit": "ns/op\t 2984573 B/op\t   60446 allocs/op",
            "extra": "144 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8271608,
            "unit": "ns/op",
            "extra": "144 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984573,
            "unit": "B/op",
            "extra": "144 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "144 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9685710,
            "unit": "ns/op\t15163754 B/op\t   23992 allocs/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9685710,
            "unit": "ns/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15163754,
            "unit": "B/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23992,
            "unit": "allocs/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9698673,
            "unit": "ns/op\t15164499 B/op\t   23996 allocs/op",
            "extra": "123 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9698673,
            "unit": "ns/op",
            "extra": "123 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164499,
            "unit": "B/op",
            "extra": "123 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23996,
            "unit": "allocs/op",
            "extra": "123 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1592d55925eb2b4c87dea77b77a77e3637704474",
          "message": "chore(deps): lock file maintenance",
          "timestamp": "2026-09-21T00:59:50Z",
          "tree_id": "e8679fb45e1864f07326590778544e4019522a99",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/1592d55925eb2b4c87dea77b77a77e3637704474"
        },
        "date": 1789952571464,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13874,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "81092 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13874,
            "unit": "ns/op",
            "extra": "81092 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "81092 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "81092 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 12544,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "95427 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 12544,
            "unit": "ns/op",
            "extra": "95427 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "95427 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "95427 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 941.1,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "1272667 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 941.1,
            "unit": "ns/op",
            "extra": "1272667 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1272667 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1272667 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 930.7,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "1289106 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 930.7,
            "unit": "ns/op",
            "extra": "1289106 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "1289106 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "1289106 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 231.3,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5249698 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 231.3,
            "unit": "ns/op",
            "extra": "5249698 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5249698 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5249698 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 247.7,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "4774620 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 247.7,
            "unit": "ns/op",
            "extra": "4774620 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "4774620 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "4774620 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 284906,
            "unit": "ns/op\t 230.03 MB/s\t    7605 B/op\t      93 allocs/op",
            "extra": "4168 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 284906,
            "unit": "ns/op",
            "extra": "4168 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 230.03,
            "unit": "MB/s",
            "extra": "4168 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7605,
            "unit": "B/op",
            "extra": "4168 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "4168 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 280487,
            "unit": "ns/op\t 233.65 MB/s\t    7575 B/op\t      93 allocs/op",
            "extra": "4202 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 280487,
            "unit": "ns/op",
            "extra": "4202 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 233.65,
            "unit": "MB/s",
            "extra": "4202 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7575,
            "unit": "B/op",
            "extra": "4202 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "4202 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 30404,
            "unit": "ns/op\t   33349 B/op\t       8 allocs/op",
            "extra": "39296 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 30404,
            "unit": "ns/op",
            "extra": "39296 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33349,
            "unit": "B/op",
            "extra": "39296 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "39296 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 30521,
            "unit": "ns/op\t   33349 B/op\t       8 allocs/op",
            "extra": "39158 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 30521,
            "unit": "ns/op",
            "extra": "39158 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33349,
            "unit": "B/op",
            "extra": "39158 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "39158 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 5362,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "217660 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 5362,
            "unit": "ns/op",
            "extra": "217660 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "217660 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "217660 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 5358,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "219094 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 5358,
            "unit": "ns/op",
            "extra": "219094 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "219094 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "219094 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 94736,
            "unit": "ns/op\t   17403 B/op\t     194 allocs/op",
            "extra": "12589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 94736,
            "unit": "ns/op",
            "extra": "12589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17403,
            "unit": "B/op",
            "extra": "12589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "12589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 97320,
            "unit": "ns/op\t   17395 B/op\t     194 allocs/op",
            "extra": "12404 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 97320,
            "unit": "ns/op",
            "extra": "12404 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17395,
            "unit": "B/op",
            "extra": "12404 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "12404 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 328241,
            "unit": "ns/op\t   73864 B/op\t     648 allocs/op",
            "extra": "3304 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 328241,
            "unit": "ns/op",
            "extra": "3304 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73864,
            "unit": "B/op",
            "extra": "3304 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3304 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 329141,
            "unit": "ns/op\t   73826 B/op\t     648 allocs/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 329141,
            "unit": "ns/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73826,
            "unit": "B/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3567 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 96593,
            "unit": "ns/op\t   14729 B/op\t     109 allocs/op",
            "extra": "12446 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 96593,
            "unit": "ns/op",
            "extra": "12446 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14729,
            "unit": "B/op",
            "extra": "12446 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "12446 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 97583,
            "unit": "ns/op\t   14716 B/op\t     109 allocs/op",
            "extra": "12339 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 97583,
            "unit": "ns/op",
            "extra": "12339 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14716,
            "unit": "B/op",
            "extra": "12339 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "12339 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 94744,
            "unit": "ns/op\t   13040 B/op\t     104 allocs/op",
            "extra": "12634 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 94744,
            "unit": "ns/op",
            "extra": "12634 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13040,
            "unit": "B/op",
            "extra": "12634 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "12634 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 95847,
            "unit": "ns/op\t   13027 B/op\t     104 allocs/op",
            "extra": "12506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 95847,
            "unit": "ns/op",
            "extra": "12506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13027,
            "unit": "B/op",
            "extra": "12506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "12506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 395547,
            "unit": "ns/op\t   82303 B/op\t     678 allocs/op",
            "extra": "2868 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 395547,
            "unit": "ns/op",
            "extra": "2868 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82303,
            "unit": "B/op",
            "extra": "2868 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2868 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 393955,
            "unit": "ns/op\t   82253 B/op\t     678 allocs/op",
            "extra": "3068 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 393955,
            "unit": "ns/op",
            "extra": "3068 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82253,
            "unit": "B/op",
            "extra": "3068 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "3068 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 98125,
            "unit": "ns/op\t   15555 B/op\t     224 allocs/op",
            "extra": "12121 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 98125,
            "unit": "ns/op",
            "extra": "12121 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15555,
            "unit": "B/op",
            "extra": "12121 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "12121 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 101714,
            "unit": "ns/op\t   15545 B/op\t     224 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 101714,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15545,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 187818,
            "unit": "ns/op\t   27390 B/op\t     354 allocs/op",
            "extra": "6086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 187818,
            "unit": "ns/op",
            "extra": "6086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27390,
            "unit": "B/op",
            "extra": "6086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "6086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 186765,
            "unit": "ns/op\t   27371 B/op\t     354 allocs/op",
            "extra": "6002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 186765,
            "unit": "ns/op",
            "extra": "6002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27371,
            "unit": "B/op",
            "extra": "6002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "6002 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 348345,
            "unit": "ns/op\t  947317 B/op\t    1255 allocs/op",
            "extra": "3309 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 348345,
            "unit": "ns/op",
            "extra": "3309 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947317,
            "unit": "B/op",
            "extra": "3309 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3309 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 351296,
            "unit": "ns/op\t  947406 B/op\t    1255 allocs/op",
            "extra": "3351 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 351296,
            "unit": "ns/op",
            "extra": "3351 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947406,
            "unit": "B/op",
            "extra": "3351 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3351 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 95952,
            "unit": "ns/op\t   15122 B/op\t     208 allocs/op",
            "extra": "12507 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 95952,
            "unit": "ns/op",
            "extra": "12507 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15122,
            "unit": "B/op",
            "extra": "12507 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "12507 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 94435,
            "unit": "ns/op\t   15114 B/op\t     208 allocs/op",
            "extra": "12637 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 94435,
            "unit": "ns/op",
            "extra": "12637 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15114,
            "unit": "B/op",
            "extra": "12637 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "12637 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 91936,
            "unit": "ns/op\t   14515 B/op\t     174 allocs/op",
            "extra": "13010 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 91936,
            "unit": "ns/op",
            "extra": "13010 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14515,
            "unit": "B/op",
            "extra": "13010 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "13010 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 91053,
            "unit": "ns/op\t   14506 B/op\t     174 allocs/op",
            "extra": "13072 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 91053,
            "unit": "ns/op",
            "extra": "13072 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14506,
            "unit": "B/op",
            "extra": "13072 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "13072 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 526994,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2288 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 526994,
            "unit": "ns/op",
            "extra": "2288 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2288 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2288 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 522811,
            "unit": "ns/op\t  290019 B/op\t    3969 allocs/op",
            "extra": "2271 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 522811,
            "unit": "ns/op",
            "extra": "2271 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290019,
            "unit": "B/op",
            "extra": "2271 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2271 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 329686,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3561 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 329686,
            "unit": "ns/op",
            "extra": "3561 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3561 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3561 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 332924,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3642 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 332924,
            "unit": "ns/op",
            "extra": "3642 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3642 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3642 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 889694,
            "unit": "ns/op\t  335232 B/op\t     573 allocs/op",
            "extra": "1335 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 889694,
            "unit": "ns/op",
            "extra": "1335 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335232,
            "unit": "B/op",
            "extra": "1335 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1335 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 888499,
            "unit": "ns/op\t  340373 B/op\t     573 allocs/op",
            "extra": "1326 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 888499,
            "unit": "ns/op",
            "extra": "1326 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 340373,
            "unit": "B/op",
            "extra": "1326 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1326 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44858006 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.71,
            "unit": "ns/op",
            "extra": "44858006 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44858006 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44858006 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45034339 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.65,
            "unit": "ns/op",
            "extra": "45034339 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45034339 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45034339 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 99.59,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "12364017 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 99.59,
            "unit": "ns/op",
            "extra": "12364017 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "12364017 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "12364017 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 91.33,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "13094436 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 91.33,
            "unit": "ns/op",
            "extra": "13094436 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "13094436 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "13094436 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 125.8,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9567062 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 125.8,
            "unit": "ns/op",
            "extra": "9567062 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9567062 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9567062 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.4,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9546188 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.4,
            "unit": "ns/op",
            "extra": "9546188 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9546188 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9546188 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 301.6,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3956613 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 301.6,
            "unit": "ns/op",
            "extra": "3956613 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3956613 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3956613 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 298.4,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "4031584 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 298.4,
            "unit": "ns/op",
            "extra": "4031584 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "4031584 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "4031584 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 74074,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16218 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 74074,
            "unit": "ns/op",
            "extra": "16218 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16218 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16218 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 74480,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16076 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 74480,
            "unit": "ns/op",
            "extra": "16076 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16076 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16076 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8211639,
            "unit": "ns/op\t 2984574 B/op\t   60446 allocs/op",
            "extra": "145 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8211639,
            "unit": "ns/op",
            "extra": "145 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984574,
            "unit": "B/op",
            "extra": "145 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "145 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8220443,
            "unit": "ns/op\t 2984571 B/op\t   60446 allocs/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8220443,
            "unit": "ns/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984571,
            "unit": "B/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "145 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9748076,
            "unit": "ns/op\t15163873 B/op\t   23992 allocs/op",
            "extra": "122 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9748076,
            "unit": "ns/op",
            "extra": "122 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15163873,
            "unit": "B/op",
            "extra": "122 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23992,
            "unit": "allocs/op",
            "extra": "122 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9867226,
            "unit": "ns/op\t15164186 B/op\t   23994 allocs/op",
            "extra": "120 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9867226,
            "unit": "ns/op",
            "extra": "120 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164186,
            "unit": "B/op",
            "extra": "120 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23994,
            "unit": "allocs/op",
            "extra": "120 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "74dd9580e3c691f83243967bc8e8dd681dd5c8b0",
          "message": "chore(deps): lock file maintenance (#552)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-21T01:02:35Z",
          "tree_id": "e8679fb45e1864f07326590778544e4019522a99",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/74dd9580e3c691f83243967bc8e8dd681dd5c8b0"
        },
        "date": 1789952771231,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 17201,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "61071 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 17201,
            "unit": "ns/op",
            "extra": "61071 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "61071 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "61071 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 14153,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "87698 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 14153,
            "unit": "ns/op",
            "extra": "87698 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "87698 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "87698 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2055,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "582747 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2055,
            "unit": "ns/op",
            "extra": "582747 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "582747 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "582747 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2114,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "500235 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2114,
            "unit": "ns/op",
            "extra": "500235 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "500235 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "500235 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 225.1,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5463496 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 225.1,
            "unit": "ns/op",
            "extra": "5463496 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5463496 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5463496 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 227.4,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5202902 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 227.4,
            "unit": "ns/op",
            "extra": "5202902 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5202902 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5202902 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 373502,
            "unit": "ns/op\t 175.46 MB/s\t    7719 B/op\t      93 allocs/op",
            "extra": "3099 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 373502,
            "unit": "ns/op",
            "extra": "3099 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 175.46,
            "unit": "MB/s",
            "extra": "3099 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7719,
            "unit": "B/op",
            "extra": "3099 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3099 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 370765,
            "unit": "ns/op\t 176.76 MB/s\t    7597 B/op\t      93 allocs/op",
            "extra": "3217 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 370765,
            "unit": "ns/op",
            "extra": "3217 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 176.76,
            "unit": "MB/s",
            "extra": "3217 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7597,
            "unit": "B/op",
            "extra": "3217 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3217 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 37289,
            "unit": "ns/op\t   33334 B/op\t       8 allocs/op",
            "extra": "31534 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 37289,
            "unit": "ns/op",
            "extra": "31534 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33334,
            "unit": "B/op",
            "extra": "31534 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "31534 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36439,
            "unit": "ns/op\t   33302 B/op\t       8 allocs/op",
            "extra": "33079 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36439,
            "unit": "ns/op",
            "extra": "33079 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33302,
            "unit": "B/op",
            "extra": "33079 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "33079 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10343,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "113667 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10343,
            "unit": "ns/op",
            "extra": "113667 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "113667 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "113667 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10307,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "115338 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10307,
            "unit": "ns/op",
            "extra": "115338 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "115338 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "115338 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 139636,
            "unit": "ns/op\t   17411 B/op\t     194 allocs/op",
            "extra": "8778 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 139636,
            "unit": "ns/op",
            "extra": "8778 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17411,
            "unit": "B/op",
            "extra": "8778 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8778 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 139364,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "9727 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 139364,
            "unit": "ns/op",
            "extra": "9727 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "9727 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "9727 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 446779,
            "unit": "ns/op\t   73869 B/op\t     648 allocs/op",
            "extra": "2623 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 446779,
            "unit": "ns/op",
            "extra": "2623 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73869,
            "unit": "B/op",
            "extra": "2623 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2623 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 443366,
            "unit": "ns/op\t   73816 B/op\t     648 allocs/op",
            "extra": "2672 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 443366,
            "unit": "ns/op",
            "extra": "2672 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73816,
            "unit": "B/op",
            "extra": "2672 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2672 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 138322,
            "unit": "ns/op\t   14735 B/op\t     109 allocs/op",
            "extra": "8850 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 138322,
            "unit": "ns/op",
            "extra": "8850 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14735,
            "unit": "B/op",
            "extra": "8850 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8850 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 141158,
            "unit": "ns/op\t   14716 B/op\t     109 allocs/op",
            "extra": "9472 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 141158,
            "unit": "ns/op",
            "extra": "9472 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14716,
            "unit": "B/op",
            "extra": "9472 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9472 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 136834,
            "unit": "ns/op\t   13046 B/op\t     104 allocs/op",
            "extra": "8589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 136834,
            "unit": "ns/op",
            "extra": "8589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13046,
            "unit": "B/op",
            "extra": "8589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8589 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 140426,
            "unit": "ns/op\t   13027 B/op\t     104 allocs/op",
            "extra": "8223 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 140426,
            "unit": "ns/op",
            "extra": "8223 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13027,
            "unit": "B/op",
            "extra": "8223 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8223 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 570158,
            "unit": "ns/op\t   82348 B/op\t     678 allocs/op",
            "extra": "1935 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 570158,
            "unit": "ns/op",
            "extra": "1935 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82348,
            "unit": "B/op",
            "extra": "1935 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "1935 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 565114,
            "unit": "ns/op\t   82291 B/op\t     678 allocs/op",
            "extra": "2076 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 565114,
            "unit": "ns/op",
            "extra": "2076 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82291,
            "unit": "B/op",
            "extra": "2076 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2076 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 143691,
            "unit": "ns/op\t   15560 B/op\t     224 allocs/op",
            "extra": "8336 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 143691,
            "unit": "ns/op",
            "extra": "8336 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15560,
            "unit": "B/op",
            "extra": "8336 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8336 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 143346,
            "unit": "ns/op\t   15547 B/op\t     224 allocs/op",
            "extra": "8659 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 143346,
            "unit": "ns/op",
            "extra": "8659 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15547,
            "unit": "B/op",
            "extra": "8659 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8659 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 270842,
            "unit": "ns/op\t   27396 B/op\t     354 allocs/op",
            "extra": "4622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 270842,
            "unit": "ns/op",
            "extra": "4622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27396,
            "unit": "B/op",
            "extra": "4622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 269275,
            "unit": "ns/op\t   27377 B/op\t     354 allocs/op",
            "extra": "4341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 269275,
            "unit": "ns/op",
            "extra": "4341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27377,
            "unit": "B/op",
            "extra": "4341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 387723,
            "unit": "ns/op\t  947420 B/op\t    1256 allocs/op",
            "extra": "3080 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 387723,
            "unit": "ns/op",
            "extra": "3080 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947420,
            "unit": "B/op",
            "extra": "3080 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3080 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 391322,
            "unit": "ns/op\t  947608 B/op\t    1256 allocs/op",
            "extra": "3025 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 391322,
            "unit": "ns/op",
            "extra": "3025 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947608,
            "unit": "B/op",
            "extra": "3025 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3025 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 140399,
            "unit": "ns/op\t   15132 B/op\t     208 allocs/op",
            "extra": "8316 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 140399,
            "unit": "ns/op",
            "extra": "8316 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15132,
            "unit": "B/op",
            "extra": "8316 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "8316 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 137871,
            "unit": "ns/op\t   15114 B/op\t     208 allocs/op",
            "extra": "9265 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 137871,
            "unit": "ns/op",
            "extra": "9265 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15114,
            "unit": "B/op",
            "extra": "9265 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9265 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 134836,
            "unit": "ns/op\t   14520 B/op\t     174 allocs/op",
            "extra": "8793 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 134836,
            "unit": "ns/op",
            "extra": "8793 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14520,
            "unit": "B/op",
            "extra": "8793 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "8793 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 133294,
            "unit": "ns/op\t   14507 B/op\t     174 allocs/op",
            "extra": "9506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 133294,
            "unit": "ns/op",
            "extra": "9506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14507,
            "unit": "B/op",
            "extra": "9506 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9506 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 543653,
            "unit": "ns/op\t  290018 B/op\t    3969 allocs/op",
            "extra": "2190 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 543653,
            "unit": "ns/op",
            "extra": "2190 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290018,
            "unit": "B/op",
            "extra": "2190 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2190 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 547204,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2116 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 547204,
            "unit": "ns/op",
            "extra": "2116 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2116 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2116 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 338819,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3489 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 338819,
            "unit": "ns/op",
            "extra": "3489 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3489 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3489 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 339027,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3488 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 339027,
            "unit": "ns/op",
            "extra": "3488 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3488 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3488 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 973096,
            "unit": "ns/op\t  341727 B/op\t     573 allocs/op",
            "extra": "1206 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 973096,
            "unit": "ns/op",
            "extra": "1206 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 341727,
            "unit": "B/op",
            "extra": "1206 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1206 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 951844,
            "unit": "ns/op\t  333243 B/op\t     573 allocs/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 951844,
            "unit": "ns/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 333243,
            "unit": "B/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44778901 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.73,
            "unit": "ns/op",
            "extra": "44778901 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44778901 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44778901 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44590455 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.9,
            "unit": "ns/op",
            "extra": "44590455 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44590455 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44590455 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 102.1,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11659984 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 102.1,
            "unit": "ns/op",
            "extra": "11659984 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11659984 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11659984 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 100.5,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11932573 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 100.5,
            "unit": "ns/op",
            "extra": "11932573 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11932573 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11932573 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.3,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9188968 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.3,
            "unit": "ns/op",
            "extra": "9188968 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9188968 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9188968 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 129.3,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9011456 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 129.3,
            "unit": "ns/op",
            "extra": "9011456 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9011456 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9011456 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 342.3,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3545560 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 342.3,
            "unit": "ns/op",
            "extra": "3545560 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3545560 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3545560 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 334.1,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3626625 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 334.1,
            "unit": "ns/op",
            "extra": "3626625 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3626625 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3626625 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 72673,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16516 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 72673,
            "unit": "ns/op",
            "extra": "16516 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16516 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16516 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73198,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16412 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73198,
            "unit": "ns/op",
            "extra": "16412 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16412 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16412 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9272652,
            "unit": "ns/op\t 2984572 B/op\t   60446 allocs/op",
            "extra": "128 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9272652,
            "unit": "ns/op",
            "extra": "128 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984572,
            "unit": "B/op",
            "extra": "128 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "128 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9066865,
            "unit": "ns/op\t 2984574 B/op\t   60446 allocs/op",
            "extra": "132 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9066865,
            "unit": "ns/op",
            "extra": "132 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984574,
            "unit": "B/op",
            "extra": "132 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "132 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9840781,
            "unit": "ns/op\t15164188 B/op\t   23994 allocs/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9840781,
            "unit": "ns/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164188,
            "unit": "B/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23994,
            "unit": "allocs/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9828365,
            "unit": "ns/op\t15164186 B/op\t   23994 allocs/op",
            "extra": "122 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9828365,
            "unit": "ns/op",
            "extra": "122 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164186,
            "unit": "B/op",
            "extra": "122 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23994,
            "unit": "allocs/op",
            "extra": "122 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2bfc7db53b10c86d66861f405c685ea67f7d4c13",
          "message": "chore(deps): update docker/setup-buildx-action action to v4",
          "timestamp": "2026-09-21T04:54:45Z",
          "tree_id": "164adc924af6b7ffeb3167d7abd685decf915c3e",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/2bfc7db53b10c86d66861f405c685ea67f7d4c13"
        },
        "date": 1789966675846,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 15590,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "73195 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 15590,
            "unit": "ns/op",
            "extra": "73195 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "73195 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "73195 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 14334,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "87061 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 14334,
            "unit": "ns/op",
            "extra": "87061 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "87061 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "87061 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2092,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "578361 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2092,
            "unit": "ns/op",
            "extra": "578361 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "578361 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "578361 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2110,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "551834 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2110,
            "unit": "ns/op",
            "extra": "551834 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "551834 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "551834 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 239.1,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "4514205 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 239.1,
            "unit": "ns/op",
            "extra": "4514205 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "4514205 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "4514205 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 252.9,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5081358 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 252.9,
            "unit": "ns/op",
            "extra": "5081358 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5081358 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5081358 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 354747,
            "unit": "ns/op\t 184.74 MB/s\t    7616 B/op\t      93 allocs/op",
            "extra": "3410 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 354747,
            "unit": "ns/op",
            "extra": "3410 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 184.74,
            "unit": "MB/s",
            "extra": "3410 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7616,
            "unit": "B/op",
            "extra": "3410 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3410 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 351108,
            "unit": "ns/op\t 186.65 MB/s\t    7661 B/op\t      93 allocs/op",
            "extra": "3374 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 351108,
            "unit": "ns/op",
            "extra": "3374 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 186.65,
            "unit": "MB/s",
            "extra": "3374 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7661,
            "unit": "B/op",
            "extra": "3374 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3374 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36559,
            "unit": "ns/op\t   33322 B/op\t       8 allocs/op",
            "extra": "32900 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36559,
            "unit": "ns/op",
            "extra": "32900 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33322,
            "unit": "B/op",
            "extra": "32900 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32900 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36489,
            "unit": "ns/op\t   33298 B/op\t       8 allocs/op",
            "extra": "32872 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36489,
            "unit": "ns/op",
            "extra": "32872 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33298,
            "unit": "B/op",
            "extra": "32872 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32872 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10303,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "113724 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10303,
            "unit": "ns/op",
            "extra": "113724 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "113724 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "113724 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10274,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "116371 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10274,
            "unit": "ns/op",
            "extra": "116371 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "116371 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "116371 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 127212,
            "unit": "ns/op\t   17409 B/op\t     194 allocs/op",
            "extra": "8502 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 127212,
            "unit": "ns/op",
            "extra": "8502 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17409,
            "unit": "B/op",
            "extra": "8502 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8502 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 125376,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "8914 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 125376,
            "unit": "ns/op",
            "extra": "8914 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "8914 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8914 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 422291,
            "unit": "ns/op\t   73865 B/op\t     648 allocs/op",
            "extra": "2676 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 422291,
            "unit": "ns/op",
            "extra": "2676 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73865,
            "unit": "B/op",
            "extra": "2676 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2676 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 421166,
            "unit": "ns/op\t   73819 B/op\t     648 allocs/op",
            "extra": "2808 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 421166,
            "unit": "ns/op",
            "extra": "2808 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73819,
            "unit": "B/op",
            "extra": "2808 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2808 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 127601,
            "unit": "ns/op\t   14735 B/op\t     109 allocs/op",
            "extra": "9086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 127601,
            "unit": "ns/op",
            "extra": "9086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14735,
            "unit": "B/op",
            "extra": "9086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9086 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 125300,
            "unit": "ns/op\t   14716 B/op\t     109 allocs/op",
            "extra": "9854 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 125300,
            "unit": "ns/op",
            "extra": "9854 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14716,
            "unit": "B/op",
            "extra": "9854 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9854 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 123770,
            "unit": "ns/op\t   13045 B/op\t     104 allocs/op",
            "extra": "8730 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 123770,
            "unit": "ns/op",
            "extra": "8730 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13045,
            "unit": "B/op",
            "extra": "8730 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8730 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 123983,
            "unit": "ns/op\t   13025 B/op\t     104 allocs/op",
            "extra": "9219 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 123983,
            "unit": "ns/op",
            "extra": "9219 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13025,
            "unit": "B/op",
            "extra": "9219 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "9219 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 520358,
            "unit": "ns/op\t   82321 B/op\t     678 allocs/op",
            "extra": "2250 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 520358,
            "unit": "ns/op",
            "extra": "2250 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82321,
            "unit": "B/op",
            "extra": "2250 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2250 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 522742,
            "unit": "ns/op\t   82261 B/op\t     678 allocs/op",
            "extra": "2269 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 522742,
            "unit": "ns/op",
            "extra": "2269 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82261,
            "unit": "B/op",
            "extra": "2269 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2269 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 130313,
            "unit": "ns/op\t   15559 B/op\t     224 allocs/op",
            "extra": "8869 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 130313,
            "unit": "ns/op",
            "extra": "8869 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15559,
            "unit": "B/op",
            "extra": "8869 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8869 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 126628,
            "unit": "ns/op\t   15545 B/op\t     224 allocs/op",
            "extra": "8534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 126628,
            "unit": "ns/op",
            "extra": "8534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15545,
            "unit": "B/op",
            "extra": "8534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8534 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 246877,
            "unit": "ns/op\t   27398 B/op\t     354 allocs/op",
            "extra": "4893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 246877,
            "unit": "ns/op",
            "extra": "4893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27398,
            "unit": "B/op",
            "extra": "4893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4893 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 246495,
            "unit": "ns/op\t   27370 B/op\t     354 allocs/op",
            "extra": "5139 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 246495,
            "unit": "ns/op",
            "extra": "5139 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27370,
            "unit": "B/op",
            "extra": "5139 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "5139 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 389817,
            "unit": "ns/op\t  948469 B/op\t    1256 allocs/op",
            "extra": "2983 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 389817,
            "unit": "ns/op",
            "extra": "2983 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 948469,
            "unit": "B/op",
            "extra": "2983 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "2983 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 387481,
            "unit": "ns/op\t  948483 B/op\t    1256 allocs/op",
            "extra": "3100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 387481,
            "unit": "ns/op",
            "extra": "3100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 948483,
            "unit": "B/op",
            "extra": "3100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 128352,
            "unit": "ns/op\t   15127 B/op\t     208 allocs/op",
            "extra": "9172 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 128352,
            "unit": "ns/op",
            "extra": "9172 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15127,
            "unit": "B/op",
            "extra": "9172 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9172 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 128600,
            "unit": "ns/op\t   15113 B/op\t     208 allocs/op",
            "extra": "9374 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 128600,
            "unit": "ns/op",
            "extra": "9374 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15113,
            "unit": "B/op",
            "extra": "9374 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9374 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 121975,
            "unit": "ns/op\t   14519 B/op\t     174 allocs/op",
            "extra": "9698 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 121975,
            "unit": "ns/op",
            "extra": "9698 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14519,
            "unit": "B/op",
            "extra": "9698 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9698 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 124929,
            "unit": "ns/op\t   14510 B/op\t     174 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 124929,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14510,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 544389,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2143 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 544389,
            "unit": "ns/op",
            "extra": "2143 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2143 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2143 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 548634,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2182 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 548634,
            "unit": "ns/op",
            "extra": "2182 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2182 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2182 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 341293,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 341293,
            "unit": "ns/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 341411,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3354 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 341411,
            "unit": "ns/op",
            "extra": "3354 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3354 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3354 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 960576,
            "unit": "ns/op\t  335014 B/op\t     573 allocs/op",
            "extra": "1220 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 960576,
            "unit": "ns/op",
            "extra": "1220 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335014,
            "unit": "B/op",
            "extra": "1220 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1220 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 963813,
            "unit": "ns/op\t  338176 B/op\t     573 allocs/op",
            "extra": "1210 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 963813,
            "unit": "ns/op",
            "extra": "1210 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 338176,
            "unit": "B/op",
            "extra": "1210 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1210 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44710261 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.71,
            "unit": "ns/op",
            "extra": "44710261 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44710261 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44710261 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44385894 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.9,
            "unit": "ns/op",
            "extra": "44385894 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44385894 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44385894 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 104.1,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11922518 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 104.1,
            "unit": "ns/op",
            "extra": "11922518 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11922518 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11922518 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 101.8,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11732468 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 101.8,
            "unit": "ns/op",
            "extra": "11732468 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11732468 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11732468 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.5,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9521277 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.5,
            "unit": "ns/op",
            "extra": "9521277 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9521277 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9521277 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 125.4,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9566121 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 125.4,
            "unit": "ns/op",
            "extra": "9566121 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9566121 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9566121 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 331.6,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3604915 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 331.6,
            "unit": "ns/op",
            "extra": "3604915 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3604915 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3604915 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 331.5,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3641269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 331.5,
            "unit": "ns/op",
            "extra": "3641269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3641269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3641269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 72852,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16476 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 72852,
            "unit": "ns/op",
            "extra": "16476 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16476 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16476 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73451,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16448 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73451,
            "unit": "ns/op",
            "extra": "16448 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16448 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16448 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9078993,
            "unit": "ns/op\t 2984577 B/op\t   60446 allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9078993,
            "unit": "ns/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984577,
            "unit": "B/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8959672,
            "unit": "ns/op\t 2984569 B/op\t   60446 allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8959672,
            "unit": "ns/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984569,
            "unit": "B/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9132023,
            "unit": "ns/op\t15163726 B/op\t   23991 allocs/op",
            "extra": "130 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9132023,
            "unit": "ns/op",
            "extra": "130 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15163726,
            "unit": "B/op",
            "extra": "130 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23991,
            "unit": "allocs/op",
            "extra": "130 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9109156,
            "unit": "ns/op\t15163914 B/op\t   23992 allocs/op",
            "extra": "130 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9109156,
            "unit": "ns/op",
            "extra": "130 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15163914,
            "unit": "B/op",
            "extra": "130 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23992,
            "unit": "allocs/op",
            "extra": "130 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "336416f4c8c1b025c71c8477fdbd283c81994121",
          "message": "chore(deps): update docker/setup-qemu-action action to v4",
          "timestamp": "2026-09-21T04:54:49Z",
          "tree_id": "829602db3f5a0dd69cdfc254087d9dbf2c48bba2",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/336416f4c8c1b025c71c8477fdbd283c81994121"
        },
        "date": 1789966872902,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 18167,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "58635 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 18167,
            "unit": "ns/op",
            "extra": "58635 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "58635 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "58635 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13760,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "83065 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13760,
            "unit": "ns/op",
            "extra": "83065 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "83065 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "83065 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2130,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "592389 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2130,
            "unit": "ns/op",
            "extra": "592389 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "592389 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "592389 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2153,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "582243 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2153,
            "unit": "ns/op",
            "extra": "582243 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "582243 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "582243 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 223.8,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5389502 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 223.8,
            "unit": "ns/op",
            "extra": "5389502 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5389502 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5389502 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 227.6,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5517296 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 227.6,
            "unit": "ns/op",
            "extra": "5517296 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5517296 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5517296 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 375141,
            "unit": "ns/op\t 174.70 MB/s\t    7676 B/op\t      93 allocs/op",
            "extra": "3094 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 375141,
            "unit": "ns/op",
            "extra": "3094 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 174.7,
            "unit": "MB/s",
            "extra": "3094 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7676,
            "unit": "B/op",
            "extra": "3094 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3094 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 356700,
            "unit": "ns/op\t 183.73 MB/s\t    7621 B/op\t      93 allocs/op",
            "extra": "3372 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 356700,
            "unit": "ns/op",
            "extra": "3372 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 183.73,
            "unit": "MB/s",
            "extra": "3372 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7621,
            "unit": "B/op",
            "extra": "3372 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3372 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36876,
            "unit": "ns/op\t   33393 B/op\t       8 allocs/op",
            "extra": "33087 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36876,
            "unit": "ns/op",
            "extra": "33087 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33393,
            "unit": "B/op",
            "extra": "33087 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "33087 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36598,
            "unit": "ns/op\t   33366 B/op\t       8 allocs/op",
            "extra": "32829 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36598,
            "unit": "ns/op",
            "extra": "32829 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33366,
            "unit": "B/op",
            "extra": "32829 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32829 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10937,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "112050 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10937,
            "unit": "ns/op",
            "extra": "112050 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "112050 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "112050 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10435,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "115510 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10435,
            "unit": "ns/op",
            "extra": "115510 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "115510 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "115510 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 138949,
            "unit": "ns/op\t   17409 B/op\t     194 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 138949,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17409,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 140022,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "8248 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 140022,
            "unit": "ns/op",
            "extra": "8248 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "8248 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8248 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 438366,
            "unit": "ns/op\t   73867 B/op\t     648 allocs/op",
            "extra": "2582 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 438366,
            "unit": "ns/op",
            "extra": "2582 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73867,
            "unit": "B/op",
            "extra": "2582 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2582 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 438467,
            "unit": "ns/op\t   73826 B/op\t     648 allocs/op",
            "extra": "2758 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 438467,
            "unit": "ns/op",
            "extra": "2758 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73826,
            "unit": "B/op",
            "extra": "2758 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2758 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 132219,
            "unit": "ns/op\t   14733 B/op\t     109 allocs/op",
            "extra": "9084 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 132219,
            "unit": "ns/op",
            "extra": "9084 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14733,
            "unit": "B/op",
            "extra": "9084 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9084 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 132073,
            "unit": "ns/op\t   14716 B/op\t     109 allocs/op",
            "extra": "9666 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 132073,
            "unit": "ns/op",
            "extra": "9666 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14716,
            "unit": "B/op",
            "extra": "9666 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9666 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 129703,
            "unit": "ns/op\t   13044 B/op\t     104 allocs/op",
            "extra": "9120 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 129703,
            "unit": "ns/op",
            "extra": "9120 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13044,
            "unit": "B/op",
            "extra": "9120 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "9120 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 129098,
            "unit": "ns/op\t   13026 B/op\t     104 allocs/op",
            "extra": "9422 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 129098,
            "unit": "ns/op",
            "extra": "9422 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13026,
            "unit": "B/op",
            "extra": "9422 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "9422 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 542698,
            "unit": "ns/op\t   82332 B/op\t     678 allocs/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 542698,
            "unit": "ns/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82332,
            "unit": "B/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 547157,
            "unit": "ns/op\t   82265 B/op\t     678 allocs/op",
            "extra": "2096 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 547157,
            "unit": "ns/op",
            "extra": "2096 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82265,
            "unit": "B/op",
            "extra": "2096 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2096 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 135724,
            "unit": "ns/op\t   15558 B/op\t     224 allocs/op",
            "extra": "8348 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 135724,
            "unit": "ns/op",
            "extra": "8348 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15558,
            "unit": "B/op",
            "extra": "8348 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8348 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 136155,
            "unit": "ns/op\t   15546 B/op\t     224 allocs/op",
            "extra": "9145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 136155,
            "unit": "ns/op",
            "extra": "9145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15546,
            "unit": "B/op",
            "extra": "9145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "9145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 259243,
            "unit": "ns/op\t   27396 B/op\t     354 allocs/op",
            "extra": "4712 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 259243,
            "unit": "ns/op",
            "extra": "4712 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27396,
            "unit": "B/op",
            "extra": "4712 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4712 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 257994,
            "unit": "ns/op\t   27369 B/op\t     354 allocs/op",
            "extra": "4540 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 257994,
            "unit": "ns/op",
            "extra": "4540 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27369,
            "unit": "B/op",
            "extra": "4540 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4540 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 378243,
            "unit": "ns/op\t  947173 B/op\t    1256 allocs/op",
            "extra": "3044 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 378243,
            "unit": "ns/op",
            "extra": "3044 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947173,
            "unit": "B/op",
            "extra": "3044 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3044 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 380641,
            "unit": "ns/op\t  947654 B/op\t    1256 allocs/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 380641,
            "unit": "ns/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947654,
            "unit": "B/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 135440,
            "unit": "ns/op\t   15127 B/op\t     208 allocs/op",
            "extra": "8979 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 135440,
            "unit": "ns/op",
            "extra": "8979 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15127,
            "unit": "B/op",
            "extra": "8979 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "8979 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 134664,
            "unit": "ns/op\t   15112 B/op\t     208 allocs/op",
            "extra": "9751 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 134664,
            "unit": "ns/op",
            "extra": "9751 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "9751 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9751 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 129075,
            "unit": "ns/op\t   14517 B/op\t     174 allocs/op",
            "extra": "9280 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 129075,
            "unit": "ns/op",
            "extra": "9280 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14517,
            "unit": "B/op",
            "extra": "9280 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9280 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 128707,
            "unit": "ns/op\t   14507 B/op\t     174 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 128707,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14507,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 547154,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 547154,
            "unit": "ns/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2170 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 551296,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2161 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 551296,
            "unit": "ns/op",
            "extra": "2161 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2161 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2161 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 339781,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3372 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 339781,
            "unit": "ns/op",
            "extra": "3372 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3372 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3372 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 343477,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3481 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 343477,
            "unit": "ns/op",
            "extra": "3481 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3481 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3481 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 961586,
            "unit": "ns/op\t  333256 B/op\t     573 allocs/op",
            "extra": "1232 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 961586,
            "unit": "ns/op",
            "extra": "1232 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 333256,
            "unit": "B/op",
            "extra": "1232 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1232 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 1025464,
            "unit": "ns/op\t  334982 B/op\t     573 allocs/op",
            "extra": "1209 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 1025464,
            "unit": "ns/op",
            "extra": "1209 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 334982,
            "unit": "B/op",
            "extra": "1209 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1209 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 27.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44838097 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 27.03,
            "unit": "ns/op",
            "extra": "44838097 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44838097 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44838097 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45029012 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.7,
            "unit": "ns/op",
            "extra": "45029012 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45029012 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45029012 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 104.2,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11689591 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11689591 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11689591 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11689591 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 102.8,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11775556 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 102.8,
            "unit": "ns/op",
            "extra": "11775556 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11775556 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11775556 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.1,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9313440 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.1,
            "unit": "ns/op",
            "extra": "9313440 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9313440 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9313440 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.1,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9218541 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.1,
            "unit": "ns/op",
            "extra": "9218541 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9218541 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9218541 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 330,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3622977 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 330,
            "unit": "ns/op",
            "extra": "3622977 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3622977 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3622977 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 333,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3628617 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 333,
            "unit": "ns/op",
            "extra": "3628617 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3628617 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3628617 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73139,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16333 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73139,
            "unit": "ns/op",
            "extra": "16333 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16333 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16333 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 72540,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16519 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 72540,
            "unit": "ns/op",
            "extra": "16519 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16519 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16519 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8733962,
            "unit": "ns/op\t 2984570 B/op\t   60446 allocs/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8733962,
            "unit": "ns/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984570,
            "unit": "B/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8757363,
            "unit": "ns/op\t 2984573 B/op\t   60446 allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8757363,
            "unit": "ns/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984573,
            "unit": "B/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9966669,
            "unit": "ns/op\t15164524 B/op\t   23996 allocs/op",
            "extra": "120 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9966669,
            "unit": "ns/op",
            "extra": "120 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164524,
            "unit": "B/op",
            "extra": "120 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23996,
            "unit": "allocs/op",
            "extra": "120 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9283562,
            "unit": "ns/op\t15164654 B/op\t   23997 allocs/op",
            "extra": "129 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9283562,
            "unit": "ns/op",
            "extra": "129 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164654,
            "unit": "B/op",
            "extra": "129 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23997,
            "unit": "allocs/op",
            "extra": "129 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "cf67ae40ad76c6ea46c2e7a612f40a6862dd964b",
          "message": "chore(deps): bump the actions group with 2 updates\n\nBumps the actions group with 2 updates: [docker/setup-qemu-action](https://github.com/docker/setup-qemu-action) and [docker/setup-buildx-action](https://github.com/docker/setup-buildx-action).\n\n\nUpdates `docker/setup-qemu-action` from 3 to 4\n- [Release notes](https://github.com/docker/setup-qemu-action/releases)\n- [Commits](https://github.com/docker/setup-qemu-action/compare/v3...v4)\n\nUpdates `docker/setup-buildx-action` from 3 to 4\n- [Release notes](https://github.com/docker/setup-buildx-action/releases)\n- [Commits](https://github.com/docker/setup-buildx-action/compare/v3...v4)\n\n---\nupdated-dependencies:\n- dependency-name: docker/setup-qemu-action\n  dependency-version: '4'\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: actions\n- dependency-name: docker/setup-buildx-action\n  dependency-version: '4'\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: actions\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>",
          "timestamp": "2026-09-23T07:07:51Z",
          "tree_id": "05c319c2482c03c78f0a16a736e6530ad02359cc",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/cf67ae40ad76c6ea46c2e7a612f40a6862dd964b"
        },
        "date": 1790147453574,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 14026,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "78727 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 14026,
            "unit": "ns/op",
            "extra": "78727 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "78727 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "78727 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 12224,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "97258 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 12224,
            "unit": "ns/op",
            "extra": "97258 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "97258 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "97258 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2396,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "490680 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2396,
            "unit": "ns/op",
            "extra": "490680 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "490680 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "490680 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2387,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "504660 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2387,
            "unit": "ns/op",
            "extra": "504660 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "504660 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "504660 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 226,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5390617 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 226,
            "unit": "ns/op",
            "extra": "5390617 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5390617 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5390617 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 226.1,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5473264 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 226.1,
            "unit": "ns/op",
            "extra": "5473264 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5473264 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5473264 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 365220,
            "unit": "ns/op\t 179.44 MB/s\t    7706 B/op\t      93 allocs/op",
            "extra": "3186 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 365220,
            "unit": "ns/op",
            "extra": "3186 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 179.44,
            "unit": "MB/s",
            "extra": "3186 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7706,
            "unit": "B/op",
            "extra": "3186 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3186 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 363533,
            "unit": "ns/op\t 180.28 MB/s\t    7549 B/op\t      93 allocs/op",
            "extra": "3333 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 363533,
            "unit": "ns/op",
            "extra": "3333 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 180.28,
            "unit": "MB/s",
            "extra": "3333 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7549,
            "unit": "B/op",
            "extra": "3333 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3333 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 40521,
            "unit": "ns/op\t   33424 B/op\t       8 allocs/op",
            "extra": "29622 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 40521,
            "unit": "ns/op",
            "extra": "29622 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33424,
            "unit": "B/op",
            "extra": "29622 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "29622 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 40894,
            "unit": "ns/op\t   33356 B/op\t       8 allocs/op",
            "extra": "28806 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 40894,
            "unit": "ns/op",
            "extra": "28806 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33356,
            "unit": "B/op",
            "extra": "28806 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "28806 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 12753,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "95042 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 12753,
            "unit": "ns/op",
            "extra": "95042 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "95042 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "95042 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 13084,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "93124 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 13084,
            "unit": "ns/op",
            "extra": "93124 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "93124 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "93124 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 106322,
            "unit": "ns/op\t   17410 B/op\t     194 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 106322,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17410,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 105906,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 105906,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 366885,
            "unit": "ns/op\t   73859 B/op\t     648 allocs/op",
            "extra": "2899 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 366885,
            "unit": "ns/op",
            "extra": "2899 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73859,
            "unit": "B/op",
            "extra": "2899 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2899 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 368408,
            "unit": "ns/op\t   73815 B/op\t     648 allocs/op",
            "extra": "3145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 368408,
            "unit": "ns/op",
            "extra": "3145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73815,
            "unit": "B/op",
            "extra": "3145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "3145 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 113583,
            "unit": "ns/op\t   14733 B/op\t     109 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 113583,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14733,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 113979,
            "unit": "ns/op\t   14717 B/op\t     109 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 113979,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14717,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 109181,
            "unit": "ns/op\t   13043 B/op\t     104 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 109181,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13043,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 109439,
            "unit": "ns/op\t   13027 B/op\t     104 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 109439,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13027,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 459528,
            "unit": "ns/op\t   82319 B/op\t     678 allocs/op",
            "extra": "2542 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 459528,
            "unit": "ns/op",
            "extra": "2542 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82319,
            "unit": "B/op",
            "extra": "2542 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2542 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 459703,
            "unit": "ns/op\t   82253 B/op\t     678 allocs/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 459703,
            "unit": "ns/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82253,
            "unit": "B/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 109878,
            "unit": "ns/op\t   15558 B/op\t     224 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 109878,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15558,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 109346,
            "unit": "ns/op\t   15545 B/op\t     224 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 109346,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15545,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 212962,
            "unit": "ns/op\t   27391 B/op\t     354 allocs/op",
            "extra": "5432 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 212962,
            "unit": "ns/op",
            "extra": "5432 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27391,
            "unit": "B/op",
            "extra": "5432 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "5432 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 212741,
            "unit": "ns/op\t   27368 B/op\t     354 allocs/op",
            "extra": "5794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 212741,
            "unit": "ns/op",
            "extra": "5794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27368,
            "unit": "B/op",
            "extra": "5794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "5794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 354380,
            "unit": "ns/op\t  944510 B/op\t    1255 allocs/op",
            "extra": "3301 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 354380,
            "unit": "ns/op",
            "extra": "3301 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 944510,
            "unit": "B/op",
            "extra": "3301 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3301 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 357472,
            "unit": "ns/op\t  944674 B/op\t    1255 allocs/op",
            "extra": "3420 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 357472,
            "unit": "ns/op",
            "extra": "3420 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 944674,
            "unit": "B/op",
            "extra": "3420 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3420 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 108039,
            "unit": "ns/op\t   15126 B/op\t     208 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 108039,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15126,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 110327,
            "unit": "ns/op\t   15114 B/op\t     208 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 110327,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15114,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 104600,
            "unit": "ns/op\t   14518 B/op\t     174 allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 104600,
            "unit": "ns/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14518,
            "unit": "B/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "10000 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 104072,
            "unit": "ns/op\t   14506 B/op\t     174 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 104072,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14506,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 517857,
            "unit": "ns/op\t  290018 B/op\t    3969 allocs/op",
            "extra": "2302 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 517857,
            "unit": "ns/op",
            "extra": "2302 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290018,
            "unit": "B/op",
            "extra": "2302 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2302 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 517437,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2290 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 517437,
            "unit": "ns/op",
            "extra": "2290 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2290 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2290 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 324911,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3643 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 324911,
            "unit": "ns/op",
            "extra": "3643 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3643 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3643 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 323714,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3621 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 323714,
            "unit": "ns/op",
            "extra": "3621 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3621 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3621 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 920144,
            "unit": "ns/op\t  340806 B/op\t     573 allocs/op",
            "extra": "1227 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 920144,
            "unit": "ns/op",
            "extra": "1227 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 340806,
            "unit": "B/op",
            "extra": "1227 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1227 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 907802,
            "unit": "ns/op\t  330930 B/op\t     572 allocs/op",
            "extra": "1257 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 907802,
            "unit": "ns/op",
            "extra": "1257 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 330930,
            "unit": "B/op",
            "extra": "1257 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 572,
            "unit": "allocs/op",
            "extra": "1257 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44634547 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.62,
            "unit": "ns/op",
            "extra": "44634547 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44634547 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44634547 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 27.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "43269542 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 27.49,
            "unit": "ns/op",
            "extra": "43269542 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "43269542 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "43269542 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 111.8,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "10646905 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 111.8,
            "unit": "ns/op",
            "extra": "10646905 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "10646905 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "10646905 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 115.1,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "10506510 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 115.1,
            "unit": "ns/op",
            "extra": "10506510 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "10506510 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "10506510 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 120.7,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9809883 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 120.7,
            "unit": "ns/op",
            "extra": "9809883 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9809883 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9809883 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 118.4,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "10218939 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 118.4,
            "unit": "ns/op",
            "extra": "10218939 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "10218939 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10218939 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 327.2,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3686077 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 327.2,
            "unit": "ns/op",
            "extra": "3686077 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3686077 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3686077 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 331,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3591219 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 331,
            "unit": "ns/op",
            "extra": "3591219 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3591219 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3591219 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 74104,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16140 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 74104,
            "unit": "ns/op",
            "extra": "16140 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16140 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16140 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73429,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16304 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73429,
            "unit": "ns/op",
            "extra": "16304 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16304 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16304 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 7828191,
            "unit": "ns/op\t 2984566 B/op\t   60446 allocs/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 7828191,
            "unit": "ns/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984566,
            "unit": "B/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "152 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 7769605,
            "unit": "ns/op\t 2984571 B/op\t   60446 allocs/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 7769605,
            "unit": "ns/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984571,
            "unit": "B/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "153 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10152295,
            "unit": "ns/op\t15166251 B/op\t   24006 allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10152295,
            "unit": "ns/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15166251,
            "unit": "B/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 24006,
            "unit": "allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 11369409,
            "unit": "ns/op\t15164283 B/op\t   23995 allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 11369409,
            "unit": "ns/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164283,
            "unit": "B/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23995,
            "unit": "allocs/op",
            "extra": "100 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "julien.noblet+github@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "distinct": true,
          "id": "382858f2220baefa47aa1afaa744786bbf2004b6",
          "message": "chore(deps): bump the actions group with 2 updates\n\nBumps the actions group with 2 updates: [docker/setup-qemu-action](https://github.com/docker/setup-qemu-action) and [docker/setup-buildx-action](https://github.com/docker/setup-buildx-action).\n\n\nUpdates `docker/setup-qemu-action` from 3 to 4\n- [Release notes](https://github.com/docker/setup-qemu-action/releases)\n- [Commits](https://github.com/docker/setup-qemu-action/compare/v3...v4)\n\nUpdates `docker/setup-buildx-action` from 3 to 4\n- [Release notes](https://github.com/docker/setup-buildx-action/releases)\n- [Commits](https://github.com/docker/setup-buildx-action/compare/v3...v4)\n\n---\nupdated-dependencies:\n- dependency-name: docker/setup-qemu-action\n  dependency-version: '4'\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: actions\n- dependency-name: docker/setup-buildx-action\n  dependency-version: '4'\n  dependency-type: direct:production\n  update-type: version-update:semver-major\n  dependency-group: actions\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>",
          "timestamp": "2026-09-25T00:57:44+02:00",
          "tree_id": "05c319c2482c03c78f0a16a736e6530ad02359cc",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/382858f2220baefa47aa1afaa744786bbf2004b6"
        },
        "date": 1790290865812,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 15689,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "71197 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 15689,
            "unit": "ns/op",
            "extra": "71197 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "71197 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "71197 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13550,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "87204 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13550,
            "unit": "ns/op",
            "extra": "87204 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "87204 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "87204 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2063,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "583254 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2063,
            "unit": "ns/op",
            "extra": "583254 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "583254 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "583254 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2051,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "591422 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2051,
            "unit": "ns/op",
            "extra": "591422 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "591422 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "591422 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 215.5,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5362680 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 215.5,
            "unit": "ns/op",
            "extra": "5362680 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5362680 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5362680 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 216.3,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5565044 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 216.3,
            "unit": "ns/op",
            "extra": "5565044 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5565044 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5565044 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 366736,
            "unit": "ns/op\t 178.70 MB/s\t    7616 B/op\t      93 allocs/op",
            "extra": "3170 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 366736,
            "unit": "ns/op",
            "extra": "3170 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 178.7,
            "unit": "MB/s",
            "extra": "3170 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7616,
            "unit": "B/op",
            "extra": "3170 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3170 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 365068,
            "unit": "ns/op\t 179.52 MB/s\t    7636 B/op\t      93 allocs/op",
            "extra": "3187 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 365068,
            "unit": "ns/op",
            "extra": "3187 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 179.52,
            "unit": "MB/s",
            "extra": "3187 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7636,
            "unit": "B/op",
            "extra": "3187 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3187 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36405,
            "unit": "ns/op\t   33310 B/op\t       8 allocs/op",
            "extra": "32880 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36405,
            "unit": "ns/op",
            "extra": "32880 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33310,
            "unit": "B/op",
            "extra": "32880 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32880 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36403,
            "unit": "ns/op\t   33294 B/op\t       8 allocs/op",
            "extra": "32961 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36403,
            "unit": "ns/op",
            "extra": "32961 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33294,
            "unit": "B/op",
            "extra": "32961 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32961 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10369,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "114381 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10369,
            "unit": "ns/op",
            "extra": "114381 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "114381 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "114381 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10352,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "116335 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10352,
            "unit": "ns/op",
            "extra": "116335 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "116335 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "116335 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 138926,
            "unit": "ns/op\t   17408 B/op\t     194 allocs/op",
            "extra": "8932 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 138926,
            "unit": "ns/op",
            "extra": "8932 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17408,
            "unit": "B/op",
            "extra": "8932 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "8932 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 138457,
            "unit": "ns/op\t   17395 B/op\t     194 allocs/op",
            "extra": "9004 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 138457,
            "unit": "ns/op",
            "extra": "9004 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17395,
            "unit": "B/op",
            "extra": "9004 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "9004 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 445449,
            "unit": "ns/op\t   73870 B/op\t     648 allocs/op",
            "extra": "2588 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 445449,
            "unit": "ns/op",
            "extra": "2588 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73870,
            "unit": "B/op",
            "extra": "2588 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2588 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 443557,
            "unit": "ns/op\t   73819 B/op\t     648 allocs/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 443557,
            "unit": "ns/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73819,
            "unit": "B/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2674 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 138645,
            "unit": "ns/op\t   14737 B/op\t     109 allocs/op",
            "extra": "8326 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 138645,
            "unit": "ns/op",
            "extra": "8326 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14737,
            "unit": "B/op",
            "extra": "8326 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8326 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 138255,
            "unit": "ns/op\t   14716 B/op\t     109 allocs/op",
            "extra": "8820 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 138255,
            "unit": "ns/op",
            "extra": "8820 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14716,
            "unit": "B/op",
            "extra": "8820 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8820 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 134886,
            "unit": "ns/op\t   13046 B/op\t     104 allocs/op",
            "extra": "8554 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 134886,
            "unit": "ns/op",
            "extra": "8554 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13046,
            "unit": "B/op",
            "extra": "8554 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8554 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 135945,
            "unit": "ns/op\t   13036 B/op\t     104 allocs/op",
            "extra": "8353 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 135945,
            "unit": "ns/op",
            "extra": "8353 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13036,
            "unit": "B/op",
            "extra": "8353 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8353 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 549329,
            "unit": "ns/op\t   82318 B/op\t     678 allocs/op",
            "extra": "2146 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 549329,
            "unit": "ns/op",
            "extra": "2146 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82318,
            "unit": "B/op",
            "extra": "2146 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2146 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 551841,
            "unit": "ns/op\t   82265 B/op\t     678 allocs/op",
            "extra": "2115 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 551841,
            "unit": "ns/op",
            "extra": "2115 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82265,
            "unit": "B/op",
            "extra": "2115 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2115 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 141989,
            "unit": "ns/op\t   15559 B/op\t     224 allocs/op",
            "extra": "8408 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 141989,
            "unit": "ns/op",
            "extra": "8408 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15559,
            "unit": "B/op",
            "extra": "8408 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8408 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 140605,
            "unit": "ns/op\t   15548 B/op\t     224 allocs/op",
            "extra": "8505 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 140605,
            "unit": "ns/op",
            "extra": "8505 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15548,
            "unit": "B/op",
            "extra": "8505 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8505 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 269726,
            "unit": "ns/op\t   27397 B/op\t     354 allocs/op",
            "extra": "4436 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 269726,
            "unit": "ns/op",
            "extra": "4436 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27397,
            "unit": "B/op",
            "extra": "4436 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4436 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 273695,
            "unit": "ns/op\t   27369 B/op\t     354 allocs/op",
            "extra": "4182 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 273695,
            "unit": "ns/op",
            "extra": "4182 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27369,
            "unit": "B/op",
            "extra": "4182 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4182 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 383737,
            "unit": "ns/op\t  947365 B/op\t    1256 allocs/op",
            "extra": "3026 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 383737,
            "unit": "ns/op",
            "extra": "3026 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947365,
            "unit": "B/op",
            "extra": "3026 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3026 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 383211,
            "unit": "ns/op\t  947170 B/op\t    1255 allocs/op",
            "extra": "3078 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 383211,
            "unit": "ns/op",
            "extra": "3078 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947170,
            "unit": "B/op",
            "extra": "3078 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3078 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 138260,
            "unit": "ns/op\t   15127 B/op\t     208 allocs/op",
            "extra": "8622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 138260,
            "unit": "ns/op",
            "extra": "8622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15127,
            "unit": "B/op",
            "extra": "8622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "8622 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 139310,
            "unit": "ns/op\t   15115 B/op\t     208 allocs/op",
            "extra": "8996 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 139310,
            "unit": "ns/op",
            "extra": "8996 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15115,
            "unit": "B/op",
            "extra": "8996 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "8996 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 134083,
            "unit": "ns/op\t   14518 B/op\t     174 allocs/op",
            "extra": "9168 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 134083,
            "unit": "ns/op",
            "extra": "9168 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14518,
            "unit": "B/op",
            "extra": "9168 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9168 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 134384,
            "unit": "ns/op\t   14506 B/op\t     174 allocs/op",
            "extra": "9393 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 134384,
            "unit": "ns/op",
            "extra": "9393 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14506,
            "unit": "B/op",
            "extra": "9393 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9393 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 548118,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2154 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 548118,
            "unit": "ns/op",
            "extra": "2154 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2154 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2154 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 558802,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2142 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 558802,
            "unit": "ns/op",
            "extra": "2142 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2142 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2142 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 346629,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3418 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 346629,
            "unit": "ns/op",
            "extra": "3418 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3418 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3418 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 348928,
            "unit": "ns/op\t  122584 B/op\t    3363 allocs/op",
            "extra": "3421 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 348928,
            "unit": "ns/op",
            "extra": "3421 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122584,
            "unit": "B/op",
            "extra": "3421 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3421 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 975889,
            "unit": "ns/op\t  339561 B/op\t     573 allocs/op",
            "extra": "1204 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 975889,
            "unit": "ns/op",
            "extra": "1204 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 339561,
            "unit": "B/op",
            "extra": "1204 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1204 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 963025,
            "unit": "ns/op\t  336717 B/op\t     573 allocs/op",
            "extra": "1207 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 963025,
            "unit": "ns/op",
            "extra": "1207 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 336717,
            "unit": "B/op",
            "extra": "1207 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1207 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44175529 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.87,
            "unit": "ns/op",
            "extra": "44175529 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44175529 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44175529 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 30.13,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "39321033 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 30.13,
            "unit": "ns/op",
            "extra": "39321033 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "39321033 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "39321033 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 100.7,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11934540 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 100.7,
            "unit": "ns/op",
            "extra": "11934540 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11934540 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11934540 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 102,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11831269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 102,
            "unit": "ns/op",
            "extra": "11831269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11831269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11831269 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 130.6,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9583656 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 130.6,
            "unit": "ns/op",
            "extra": "9583656 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9583656 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9583656 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 127.3,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9599088 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 127.3,
            "unit": "ns/op",
            "extra": "9599088 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9599088 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9599088 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 334.3,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3644874 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 334.3,
            "unit": "ns/op",
            "extra": "3644874 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3644874 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3644874 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 331.1,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3638217 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 331.1,
            "unit": "ns/op",
            "extra": "3638217 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3638217 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3638217 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73047,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16395 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73047,
            "unit": "ns/op",
            "extra": "16395 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16395 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16395 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73217,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16381 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73217,
            "unit": "ns/op",
            "extra": "16381 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16381 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16381 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8993160,
            "unit": "ns/op\t 2984573 B/op\t   60446 allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8993160,
            "unit": "ns/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984573,
            "unit": "B/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8965822,
            "unit": "ns/op\t 2984577 B/op\t   60446 allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8965822,
            "unit": "ns/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984577,
            "unit": "B/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10490010,
            "unit": "ns/op\t15164425 B/op\t   23995 allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10490010,
            "unit": "ns/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164425,
            "unit": "B/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23995,
            "unit": "allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10605570,
            "unit": "ns/op\t15165430 B/op\t   24001 allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10605570,
            "unit": "ns/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15165430,
            "unit": "B/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 24001,
            "unit": "allocs/op",
            "extra": "100 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "julien.noblet@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "committer": {
            "email": "julien.noblet@gmail.com",
            "name": "Julien Noblet",
            "username": "julien-noblet"
          },
          "distinct": true,
          "id": "3fd71150778cfb93a41c91769711953b83e824e3",
          "message": "fix(nix): update vendorHash in flake.nix",
          "timestamp": "2026-09-25T01:08:06+02:00",
          "tree_id": "95f8d5bac92685e3ccfa802641a8353575739b47",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/3fd71150778cfb93a41c91769711953b83e824e3"
        },
        "date": 1790291478461,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 16190,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "66175 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 16190,
            "unit": "ns/op",
            "extra": "66175 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "66175 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "66175 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13500,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "87417 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13500,
            "unit": "ns/op",
            "extra": "87417 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "87417 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "87417 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2048,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "584187 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2048,
            "unit": "ns/op",
            "extra": "584187 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "584187 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "584187 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2053,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "538584 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2053,
            "unit": "ns/op",
            "extra": "538584 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "538584 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "538584 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 225.8,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5468092 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 225.8,
            "unit": "ns/op",
            "extra": "5468092 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5468092 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5468092 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 227,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5579450 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 227,
            "unit": "ns/op",
            "extra": "5579450 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5579450 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5579450 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 365785,
            "unit": "ns/op\t 179.17 MB/s\t    7701 B/op\t      93 allocs/op",
            "extra": "3283 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 365785,
            "unit": "ns/op",
            "extra": "3283 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 179.17,
            "unit": "MB/s",
            "extra": "3283 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7701,
            "unit": "B/op",
            "extra": "3283 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3283 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 364143,
            "unit": "ns/op\t 179.97 MB/s\t    7582 B/op\t      93 allocs/op",
            "extra": "3162 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 364143,
            "unit": "ns/op",
            "extra": "3162 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 179.97,
            "unit": "MB/s",
            "extra": "3162 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7582,
            "unit": "B/op",
            "extra": "3162 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3162 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36318,
            "unit": "ns/op\t   33326 B/op\t       8 allocs/op",
            "extra": "32954 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36318,
            "unit": "ns/op",
            "extra": "32954 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33326,
            "unit": "B/op",
            "extra": "32954 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32954 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36567,
            "unit": "ns/op\t   33427 B/op\t       8 allocs/op",
            "extra": "32638 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36567,
            "unit": "ns/op",
            "extra": "32638 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33427,
            "unit": "B/op",
            "extra": "32638 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "32638 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10211,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "114771 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10211,
            "unit": "ns/op",
            "extra": "114771 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "114771 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "114771 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10206,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "117255 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10206,
            "unit": "ns/op",
            "extra": "117255 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "117255 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "117255 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 137346,
            "unit": "ns/op\t   17407 B/op\t     194 allocs/op",
            "extra": "9218 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 137346,
            "unit": "ns/op",
            "extra": "9218 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17407,
            "unit": "B/op",
            "extra": "9218 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "9218 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 136007,
            "unit": "ns/op\t   17394 B/op\t     194 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 136007,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17394,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 441951,
            "unit": "ns/op\t   73866 B/op\t     648 allocs/op",
            "extra": "2682 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 441951,
            "unit": "ns/op",
            "extra": "2682 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73866,
            "unit": "B/op",
            "extra": "2682 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2682 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 440962,
            "unit": "ns/op\t   73813 B/op\t     648 allocs/op",
            "extra": "2763 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 440962,
            "unit": "ns/op",
            "extra": "2763 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73813,
            "unit": "B/op",
            "extra": "2763 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2763 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 136953,
            "unit": "ns/op\t   14733 B/op\t     109 allocs/op",
            "extra": "8988 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 136953,
            "unit": "ns/op",
            "extra": "8988 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14733,
            "unit": "B/op",
            "extra": "8988 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8988 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 138015,
            "unit": "ns/op\t   14715 B/op\t     109 allocs/op",
            "extra": "8846 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 138015,
            "unit": "ns/op",
            "extra": "8846 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14715,
            "unit": "B/op",
            "extra": "8846 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "8846 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 132507,
            "unit": "ns/op\t   13044 B/op\t     104 allocs/op",
            "extra": "8587 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 132507,
            "unit": "ns/op",
            "extra": "8587 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13044,
            "unit": "B/op",
            "extra": "8587 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8587 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 133221,
            "unit": "ns/op\t   13032 B/op\t     104 allocs/op",
            "extra": "8948 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 133221,
            "unit": "ns/op",
            "extra": "8948 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13032,
            "unit": "B/op",
            "extra": "8948 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "8948 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 553848,
            "unit": "ns/op\t   82321 B/op\t     678 allocs/op",
            "extra": "2107 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 553848,
            "unit": "ns/op",
            "extra": "2107 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82321,
            "unit": "B/op",
            "extra": "2107 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2107 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 554123,
            "unit": "ns/op\t   82262 B/op\t     678 allocs/op",
            "extra": "2100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 554123,
            "unit": "ns/op",
            "extra": "2100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82262,
            "unit": "B/op",
            "extra": "2100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2100 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 139807,
            "unit": "ns/op\t   15560 B/op\t     224 allocs/op",
            "extra": "8860 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 139807,
            "unit": "ns/op",
            "extra": "8860 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15560,
            "unit": "B/op",
            "extra": "8860 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "8860 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 140083,
            "unit": "ns/op\t   15551 B/op\t     224 allocs/op",
            "extra": "7293 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 140083,
            "unit": "ns/op",
            "extra": "7293 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15551,
            "unit": "B/op",
            "extra": "7293 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "7293 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 263902,
            "unit": "ns/op\t   27396 B/op\t     354 allocs/op",
            "extra": "4584 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 263902,
            "unit": "ns/op",
            "extra": "4584 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27396,
            "unit": "B/op",
            "extra": "4584 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4584 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 264113,
            "unit": "ns/op\t   27369 B/op\t     354 allocs/op",
            "extra": "4382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 264113,
            "unit": "ns/op",
            "extra": "4382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27369,
            "unit": "B/op",
            "extra": "4382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 377947,
            "unit": "ns/op\t  946734 B/op\t    1255 allocs/op",
            "extra": "2934 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 377947,
            "unit": "ns/op",
            "extra": "2934 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 946734,
            "unit": "B/op",
            "extra": "2934 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "2934 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 380731,
            "unit": "ns/op\t  947074 B/op\t    1255 allocs/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 380731,
            "unit": "ns/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 947074,
            "unit": "B/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1255,
            "unit": "allocs/op",
            "extra": "3154 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 138301,
            "unit": "ns/op\t   15125 B/op\t     208 allocs/op",
            "extra": "9169 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 138301,
            "unit": "ns/op",
            "extra": "9169 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15125,
            "unit": "B/op",
            "extra": "9169 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9169 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 136587,
            "unit": "ns/op\t   15113 B/op\t     208 allocs/op",
            "extra": "9441 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 136587,
            "unit": "ns/op",
            "extra": "9441 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15113,
            "unit": "B/op",
            "extra": "9441 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9441 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 130442,
            "unit": "ns/op\t   14519 B/op\t     174 allocs/op",
            "extra": "9439 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 130442,
            "unit": "ns/op",
            "extra": "9439 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14519,
            "unit": "B/op",
            "extra": "9439 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9439 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 131713,
            "unit": "ns/op\t   14507 B/op\t     174 allocs/op",
            "extra": "9760 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 131713,
            "unit": "ns/op",
            "extra": "9760 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14507,
            "unit": "B/op",
            "extra": "9760 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9760 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 543875,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2180 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 543875,
            "unit": "ns/op",
            "extra": "2180 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2180 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2180 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 550533,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2138 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 550533,
            "unit": "ns/op",
            "extra": "2138 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2138 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2138 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 338773,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3546 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 338773,
            "unit": "ns/op",
            "extra": "3546 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3546 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3546 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 341134,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3494 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 341134,
            "unit": "ns/op",
            "extra": "3494 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3494 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3494 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 954413,
            "unit": "ns/op\t  335424 B/op\t     573 allocs/op",
            "extra": "1237 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 954413,
            "unit": "ns/op",
            "extra": "1237 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335424,
            "unit": "B/op",
            "extra": "1237 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1237 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 950758,
            "unit": "ns/op\t  333184 B/op\t     573 allocs/op",
            "extra": "1230 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 950758,
            "unit": "ns/op",
            "extra": "1230 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 333184,
            "unit": "B/op",
            "extra": "1230 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1230 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.79,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44370073 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.79,
            "unit": "ns/op",
            "extra": "44370073 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44370073 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44370073 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44915575 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.71,
            "unit": "ns/op",
            "extra": "44915575 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44915575 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44915575 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 101.2,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11642589 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 101.2,
            "unit": "ns/op",
            "extra": "11642589 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11642589 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11642589 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 106.1,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11753091 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 106.1,
            "unit": "ns/op",
            "extra": "11753091 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11753091 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11753091 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 123.8,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9618864 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 123.8,
            "unit": "ns/op",
            "extra": "9618864 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9618864 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9618864 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 126.5,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9705926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 126.5,
            "unit": "ns/op",
            "extra": "9705926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9705926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9705926 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 329.1,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3616464 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 329.1,
            "unit": "ns/op",
            "extra": "3616464 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3616464 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3616464 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 328.8,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3636880 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 328.8,
            "unit": "ns/op",
            "extra": "3636880 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3636880 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3636880 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73056,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16429 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73056,
            "unit": "ns/op",
            "extra": "16429 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16429 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16429 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 73455,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16513 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 73455,
            "unit": "ns/op",
            "extra": "16513 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16513 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16513 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8986880,
            "unit": "ns/op\t 2984573 B/op\t   60446 allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8986880,
            "unit": "ns/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984573,
            "unit": "B/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8955019,
            "unit": "ns/op\t 2984575 B/op\t   60446 allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8955019,
            "unit": "ns/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984575,
            "unit": "B/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "133 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10460165,
            "unit": "ns/op\t15164209 B/op\t   23994 allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10460165,
            "unit": "ns/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164209,
            "unit": "B/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23994,
            "unit": "allocs/op",
            "extra": "100 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10407870,
            "unit": "ns/op\t15165554 B/op\t   24002 allocs/op",
            "extra": "112 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10407870,
            "unit": "ns/op",
            "extra": "112 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15165554,
            "unit": "B/op",
            "extra": "112 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 24002,
            "unit": "allocs/op",
            "extra": "112 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "29139614+renovate[bot]@users.noreply.github.com",
            "name": "renovate[bot]",
            "username": "renovate[bot]"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9a9b43c79887f24035a5d41f5fa63cb33f5018b7",
          "message": "fix(deps): update module github.com/mark3labs/mcp-go to v1.1.1 (#555)\n\n* fix(deps): update module github.com/mark3labs/mcp-go to v1.1.1\n\n* fix(nix): update vendorHash in flake.nix\n\n* chore(deps): tidy go.sum\n\n---------\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>\nCo-authored-by: Julien Noblet <julien.noblet@gmail.com>",
          "timestamp": "2026-09-24T23:14:33Z",
          "tree_id": "dea22631e0c764c99db9137e135a17bcf21e6b31",
          "url": "https://github.com/julien-noblet/download-geofabrik/commit/9a9b43c79887f24035a5d41f5fa63cb33f5018b7"
        },
        "date": 1790291890077,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 16024,
            "unit": "ns/op\t    8075 B/op\t     125 allocs/op",
            "extra": "67087 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 16024,
            "unit": "ns/op",
            "extra": "67087 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8075,
            "unit": "B/op",
            "extra": "67087 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "67087 times"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli)",
            "value": 13988,
            "unit": "ns/op\t    8074 B/op\t     125 allocs/op",
            "extra": "90952 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - ns/op",
            "value": 13988,
            "unit": "ns/op",
            "extra": "90952 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - B/op",
            "value": 8074,
            "unit": "B/op",
            "extra": "90952 times\n2 procs"
          },
          {
            "name": "BenchmarkCLIExecuteHelp (github.com/julien-noblet/download-geofabrik/internal/cli) - allocs/op",
            "value": 125,
            "unit": "allocs/op",
            "extra": "90952 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2008,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "601971 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2008,
            "unit": "ns/op",
            "extra": "601971 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "601971 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "601971 times"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 2015,
            "unit": "ns/op\t     288 B/op\t       2 allocs/op",
            "extra": "597068 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 2015,
            "unit": "ns/op",
            "extra": "597068 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "597068 times\n2 procs"
          },
          {
            "name": "BenchmarkFileExists (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "597068 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 223.3,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "4972294 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 223.3,
            "unit": "ns/op",
            "extra": "4972294 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "4972294 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "4972294 times"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 209.7,
            "unit": "ns/op\t     800 B/op\t       6 allocs/op",
            "extra": "5953718 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 209.7,
            "unit": "ns/op",
            "extra": "5953718 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 800,
            "unit": "B/op",
            "extra": "5953718 times\n2 procs"
          },
          {
            "name": "BenchmarkNewDownloader (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "5953718 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 342452,
            "unit": "ns/op\t 191.37 MB/s\t    7687 B/op\t      93 allocs/op",
            "extra": "3471 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 342452,
            "unit": "ns/op",
            "extra": "3471 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 191.37,
            "unit": "MB/s",
            "extra": "3471 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7687,
            "unit": "B/op",
            "extra": "3471 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3471 times"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 341845,
            "unit": "ns/op\t 191.71 MB/s\t    7609 B/op\t      93 allocs/op",
            "extra": "3511 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 341845,
            "unit": "ns/op",
            "extra": "3511 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - MB/s",
            "value": 191.71,
            "unit": "MB/s",
            "extra": "3511 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 7609,
            "unit": "B/op",
            "extra": "3511 times\n2 procs"
          },
          {
            "name": "BenchmarkDownloadFileStreamMD5 (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 93,
            "unit": "allocs/op",
            "extra": "3511 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 36004,
            "unit": "ns/op\t   33337 B/op\t       8 allocs/op",
            "extra": "33274 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 36004,
            "unit": "ns/op",
            "extra": "33274 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33337,
            "unit": "B/op",
            "extra": "33274 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "33274 times"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 35987,
            "unit": "ns/op\t   33345 B/op\t       8 allocs/op",
            "extra": "33296 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 35987,
            "unit": "ns/op",
            "extra": "33296 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 33345,
            "unit": "B/op",
            "extra": "33296 times\n2 procs"
          },
          {
            "name": "BenchmarkComputeMD5Hash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "33296 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10255,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "108372 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10255,
            "unit": "ns/op",
            "extra": "108372 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "108372 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "108372 times"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader)",
            "value": 10097,
            "unit": "ns/op\t    1256 B/op\t      10 allocs/op",
            "extra": "118016 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - ns/op",
            "value": 10097,
            "unit": "ns/op",
            "extra": "118016 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - B/op",
            "value": 1256,
            "unit": "B/op",
            "extra": "118016 times\n2 procs"
          },
          {
            "name": "BenchmarkCheckFileHash (github.com/julien-noblet/download-geofabrik/internal/downloader) - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "118016 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 124508,
            "unit": "ns/op\t   17407 B/op\t     194 allocs/op",
            "extra": "9405 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 124508,
            "unit": "ns/op",
            "extra": "9405 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17407,
            "unit": "B/op",
            "extra": "9405 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "9405 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike)",
            "value": 124326,
            "unit": "ns/op\t   17393 B/op\t     194 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - ns/op",
            "value": 124326,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - B/op",
            "value": 17393,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/bbbike) - allocs/op",
            "value": 194,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 406644,
            "unit": "ns/op\t   73862 B/op\t     648 allocs/op",
            "extra": "2785 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 406644,
            "unit": "ns/op",
            "extra": "2785 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73862,
            "unit": "B/op",
            "extra": "2785 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2785 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day)",
            "value": 411191,
            "unit": "ns/op\t   73817 B/op\t     648 allocs/op",
            "extra": "2794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - ns/op",
            "value": 411191,
            "unit": "ns/op",
            "extra": "2794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - B/op",
            "value": 73817,
            "unit": "B/op",
            "extra": "2794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geo2day) - allocs/op",
            "value": 648,
            "unit": "allocs/op",
            "extra": "2794 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 122714,
            "unit": "ns/op\t   14734 B/op\t     109 allocs/op",
            "extra": "9038 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 122714,
            "unit": "ns/op",
            "extra": "9038 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14734,
            "unit": "B/op",
            "extra": "9038 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "9038 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik)",
            "value": 122910,
            "unit": "ns/op\t   14717 B/op\t     109 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - ns/op",
            "value": 122910,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - B/op",
            "value": 14717,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik) - allocs/op",
            "value": 109,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 118450,
            "unit": "ns/op\t   13044 B/op\t     104 allocs/op",
            "extra": "9284 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 118450,
            "unit": "ns/op",
            "extra": "9284 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13044,
            "unit": "B/op",
            "extra": "9284 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "9284 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda)",
            "value": 119692,
            "unit": "ns/op\t   13028 B/op\t     104 allocs/op",
            "extra": "9703 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - ns/op",
            "value": 119692,
            "unit": "ns/op",
            "extra": "9703 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - B/op",
            "value": 13028,
            "unit": "B/op",
            "extra": "9703 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/movisda) - allocs/op",
            "value": 104,
            "unit": "allocs/op",
            "extra": "9703 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 499651,
            "unit": "ns/op\t   82325 B/op\t     678 allocs/op",
            "extra": "2373 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 499651,
            "unit": "ns/op",
            "extra": "2373 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82325,
            "unit": "B/op",
            "extra": "2373 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2373 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr)",
            "value": 497291,
            "unit": "ns/op\t   82256 B/op\t     678 allocs/op",
            "extra": "2341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - ns/op",
            "value": 497291,
            "unit": "ns/op",
            "extra": "2341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - B/op",
            "value": 82256,
            "unit": "B/op",
            "extra": "2341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr) - allocs/op",
            "value": 678,
            "unit": "allocs/op",
            "extra": "2341 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 126788,
            "unit": "ns/op\t   15558 B/op\t     224 allocs/op",
            "extra": "9186 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 126788,
            "unit": "ns/op",
            "extra": "9186 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15558,
            "unit": "B/op",
            "extra": "9186 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "9186 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch)",
            "value": 128981,
            "unit": "ns/op\t   15546 B/op\t     224 allocs/op",
            "extra": "9193 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - ns/op",
            "value": 128981,
            "unit": "ns/op",
            "extra": "9193 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - B/op",
            "value": 15546,
            "unit": "B/op",
            "extra": "9193 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmch) - allocs/op",
            "value": 224,
            "unit": "allocs/op",
            "extra": "9193 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 241273,
            "unit": "ns/op\t   27392 B/op\t     354 allocs/op",
            "extra": "4824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 241273,
            "unit": "ns/op",
            "extra": "4824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27392,
            "unit": "B/op",
            "extra": "4824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "4824 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr)",
            "value": 240003,
            "unit": "ns/op\t   27373 B/op\t     354 allocs/op",
            "extra": "5046 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - ns/op",
            "value": 240003,
            "unit": "ns/op",
            "extra": "5046 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - B/op",
            "value": 27373,
            "unit": "B/op",
            "extra": "5046 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr) - allocs/op",
            "value": 354,
            "unit": "allocs/op",
            "extra": "5046 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 376163,
            "unit": "ns/op\t  948826 B/op\t    1256 allocs/op",
            "extra": "3018 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 376163,
            "unit": "ns/op",
            "extra": "3018 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 948826,
            "unit": "B/op",
            "extra": "3018 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3018 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit)",
            "value": 391391,
            "unit": "ns/op\t  948261 B/op\t    1256 allocs/op",
            "extra": "3033 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - ns/op",
            "value": 391391,
            "unit": "ns/op",
            "extra": "3033 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - B/op",
            "value": 948261,
            "unit": "B/op",
            "extra": "3033 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmit) - allocs/op",
            "value": 1256,
            "unit": "allocs/op",
            "extra": "3033 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 125140,
            "unit": "ns/op\t   15124 B/op\t     208 allocs/op",
            "extra": "9798 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 125140,
            "unit": "ns/op",
            "extra": "9798 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15124,
            "unit": "B/op",
            "extra": "9798 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9798 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu)",
            "value": 124231,
            "unit": "ns/op\t   15114 B/op\t     208 allocs/op",
            "extra": "9382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - ns/op",
            "value": 124231,
            "unit": "ns/op",
            "extra": "9382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - B/op",
            "value": 15114,
            "unit": "B/op",
            "extra": "9382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu) - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "9382 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 119152,
            "unit": "ns/op\t   14518 B/op\t     174 allocs/op",
            "extra": "9691 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 119152,
            "unit": "ns/op",
            "extra": "9691 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14518,
            "unit": "B/op",
            "extra": "9691 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9691 times"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw)",
            "value": 119645,
            "unit": "ns/op\t   14506 B/op\t     174 allocs/op",
            "extra": "9465 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - ns/op",
            "value": 119645,
            "unit": "ns/op",
            "extra": "9465 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - B/op",
            "value": 14506,
            "unit": "B/op",
            "extra": "9465 times\n2 procs"
          },
          {
            "name": "BenchmarkFetchCatalogMock (github.com/julien-noblet/download-geofabrik/internal/provider/osmtw) - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "9465 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 544775,
            "unit": "ns/op\t  290017 B/op\t    3969 allocs/op",
            "extra": "2210 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 544775,
            "unit": "ns/op",
            "extra": "2210 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290017,
            "unit": "B/op",
            "extra": "2210 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2210 times"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 546480,
            "unit": "ns/op\t  290016 B/op\t    3969 allocs/op",
            "extra": "2206 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 546480,
            "unit": "ns/op",
            "extra": "2206 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 290016,
            "unit": "B/op",
            "extra": "2206 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableStandard (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3969,
            "unit": "allocs/op",
            "extra": "2206 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 342891,
            "unit": "ns/op\t  122583 B/op\t    3363 allocs/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 342891,
            "unit": "ns/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122583,
            "unit": "B/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3486 times"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 346108,
            "unit": "ns/op\t  122582 B/op\t    3363 allocs/op",
            "extra": "3456 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 346108,
            "unit": "ns/op",
            "extra": "3456 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 122582,
            "unit": "B/op",
            "extra": "3456 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintTableMarkdown (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 3363,
            "unit": "allocs/op",
            "extra": "3456 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 966457,
            "unit": "ns/op\t  335504 B/op\t     573 allocs/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 966457,
            "unit": "ns/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335504,
            "unit": "B/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1225 times"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui)",
            "value": 964188,
            "unit": "ns/op\t  335466 B/op\t     573 allocs/op",
            "extra": "1232 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - ns/op",
            "value": 964188,
            "unit": "ns/op",
            "extra": "1232 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - B/op",
            "value": 335466,
            "unit": "B/op",
            "extra": "1232 times\n2 procs"
          },
          {
            "name": "BenchmarkPrintJSON (github.com/julien-noblet/download-geofabrik/internal/ui) - allocs/op",
            "value": 573,
            "unit": "allocs/op",
            "extra": "1232 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 27.41,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "43446226 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 27.41,
            "unit": "ns/op",
            "extra": "43446226 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "43446226 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "43446226 times"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 26.79,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44921904 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 26.79,
            "unit": "ns/op",
            "extra": "44921904 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44921904 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogExist (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44921904 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 102.8,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11710536 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 102.8,
            "unit": "ns/op",
            "extra": "11710536 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11710536 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11710536 times"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 102.5,
            "unit": "ns/op\t      96 B/op\t       1 allocs/op",
            "extra": "11486394 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 102.5,
            "unit": "ns/op",
            "extra": "11486394 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "11486394 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogGet (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "11486394 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 128.6,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9291320 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 128.6,
            "unit": "ns/op",
            "extra": "9291320 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9291320 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9291320 times"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 129.4,
            "unit": "ns/op\t     192 B/op\t       2 allocs/op",
            "extra": "9302016 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 129.4,
            "unit": "ns/op",
            "extra": "9302016 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 192,
            "unit": "B/op",
            "extra": "9302016 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogFind (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9302016 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 341.3,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3604596 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 341.3,
            "unit": "ns/op",
            "extra": "3604596 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3604596 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3604596 times"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 335.7,
            "unit": "ns/op\t     224 B/op\t       3 allocs/op",
            "extra": "3546690 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 335.7,
            "unit": "ns/op",
            "extra": "3546690 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "3546690 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogResolveURL (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "3546690 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 72867,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16435 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 72867,
            "unit": "ns/op",
            "extra": "16435 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16435 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16435 times"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 72464,
            "unit": "ns/op\t    9472 B/op\t       1 allocs/op",
            "extra": "16538 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 72464,
            "unit": "ns/op",
            "extra": "16538 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 9472,
            "unit": "B/op",
            "extra": "16538 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSortedKeys (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16538 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8733798,
            "unit": "ns/op\t 2984565 B/op\t   60446 allocs/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8733798,
            "unit": "ns/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984565,
            "unit": "B/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "136 times"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 8778412,
            "unit": "ns/op\t 2984572 B/op\t   60446 allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 8778412,
            "unit": "ns/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 2984572,
            "unit": "B/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogLoadFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 60446,
            "unit": "allocs/op",
            "extra": "136 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 9752657,
            "unit": "ns/op\t15165008 B/op\t   23999 allocs/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 9752657,
            "unit": "ns/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15165008,
            "unit": "B/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23999,
            "unit": "allocs/op",
            "extra": "123 times"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog)",
            "value": 10011043,
            "unit": "ns/op\t15164532 B/op\t   23996 allocs/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - ns/op",
            "value": 10011043,
            "unit": "ns/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - B/op",
            "value": 15164532,
            "unit": "B/op",
            "extra": "100 times\n2 procs"
          },
          {
            "name": "BenchmarkCatalogSaveFile (github.com/julien-noblet/download-geofabrik/pkg/catalog) - allocs/op",
            "value": 23996,
            "unit": "allocs/op",
            "extra": "100 times\n2 procs"
          }
        ]
      }
    ]
  }
}