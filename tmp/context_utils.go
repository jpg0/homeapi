package main

func expectedKeys(keys []string, ctx map[string]string) {

	removeBlanks(ctx)

	for _, key := range keys {
		_, present := ctx[key]

		if !present {
			ctx["!" + key] = "true"
		}
	}
}

func merge(m1, m2 map[string]string) map[string]string {
	m := make(map[string]string)

	for k, v := range m2 {
		m[k] = v
	}

	for k, v := range m1 {
		m[k] = v
	}

	return m
}

func removeBlanks(m map[string]string) {
	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}
}