jobs: {
  without_checkout: {
    steps: [
      {
        name: "Run pipeline"
        uses: =~"dagger/dagger-for-github@.+"
        with: {
          version: =~".+"
          verb: "call"
          args: =~"^check(?: .*)?"
        }
      }
    ]
  }
}
