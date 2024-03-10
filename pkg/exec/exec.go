package exec

import "context"

/**
 * @Author: root
 * @Description:
 * @File:  exec
 * @Date: 3/8/24 7:52 PM
 */

type ExecIF interface {
	Exec(ctx context.Context) error
}
