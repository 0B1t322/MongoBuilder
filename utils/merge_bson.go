package utils

import "go.mongodb.org/mongo-driver/bson"

func MergeBsonM(ms ...bson.M) bson.M {
	out := bson.M{}
	{
		if len(ms) > 1 {
			first := ms[0]
			ms = ms[1:]
			for k, v := range first {
				out[k] = v
			}
		} else if len(ms) == 1 {
			return ms[0]
		}

		with := MergeBsonM(ms...)
		for k, v := range with {
			if v, ok := v.(bson.M); ok {
				if bv, ok := out[k]; ok {
					if bv, ok := bv.(bson.M); ok {
						out[k] = MergeBsonM(bv, v)
						continue
					}
				}
			}
			out[k] = v
		}

	}
	return out
}
