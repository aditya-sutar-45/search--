package frontier

type Queue struct {
	*ReddisConnection
}

func NewQueue(r *ReddisConnection) *Queue {
	return &Queue{
		ReddisConnection: r,
	}
}

func (q *Queue) Enqueue(url string) error {
	return q.client.LPush(q.ctx, QUEUE_KEY, url).Err()
}

func (q *Queue) Dequeue() (string, error) {
	result, err := q.client.BRPop(q.ctx, 0, QUEUE_KEY).Result()
	if err != nil {
		return "", err
	}

	return result[1], nil
}
