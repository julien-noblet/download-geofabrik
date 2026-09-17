window.BENCHMARK_DATA = {
  "lastUpdate": 1789616327661,
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
      }
    ]
  }
}