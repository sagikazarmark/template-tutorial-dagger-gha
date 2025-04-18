jobs: {
  test: {
    steps: [
      {
        uses: =~"actions/checkout@v[0-9]+"
      },
      {
        name: "Run pipeline"
        uses: =~"dagger/dagger-for-github@.+"
        with: {
          version: =~".+"
          verb: "call"
          args: =~"^test(?: .*)?"
        }
      }
    ]
  }
}
