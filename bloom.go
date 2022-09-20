package redis

import "context"

type (
	BfInsertArgs struct {
		Key      string
		Values   []interface{}
		Capacity uint32
		//NoCreate bool
	}
)

func (c *Client) BfInsert(ctx context.Context, op *BfInsertArgs) *Cmd {
	args := make([]interface{}, 5, 5+len(op.Values))
	args[0] = "BF.INSERT"
	args[1] = op.Key
	cap := op.Capacity
	if cap <= 100 {
		cap = 10000
	}
	args[2] = "CAPACITY"
	args[3] = cap
	args[4] = "ITEMS"
	args = appendArgs(args, op.Values)
	return c.Do(ctx, args...)
}

func (c *Client) BfAdd(ctx context.Context, key string, value interface{}) *Cmd {

	return c.Do(ctx, "BF.ADD", key, value)

}
func (c *Client) BfMadd(ctx context.Context, key string, values ...interface{}) *Cmd {

	args := make([]interface{}, 2, 2+len(values))
	args[0] = "BF.MADD"
	args[1] = key
	args = appendArgs(args, values)
	return c.Do(ctx, args...)

}

func (c *Client) BfExists(ctx context.Context, key string, value interface{}) *Cmd {

	return c.Do(ctx, "BF.EXISTS", key, value)

}
func (c *Client) BfMexists(ctx context.Context, key string, values ...interface{}) *Cmd {

	args := make([]interface{}, 2, 2+len(values))
	args[0] = "BF.MEXISTS"
	args[1] = key
	args = appendArgs(args, values)
	return c.Do(ctx, args...)

}
