package design

import "container/heap"

/**
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
示例:

Twitter twitter = new Twitter();

// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);

// 用户1关注了用户2.
twitter.follow(1, 2);

// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);

// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);

// 用户1取消关注了用户2.
twitter.unfollow(1, 2);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-twitter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type listNode struct {
	tweet, time int
	next        *listNode
}

func newListNode(tweetId, time int, next *listNode) *listNode {
	return &listNode{tweet: tweetId, time: time, next: next}
}

type listNodeHeap []*listNode

func (l *listNodeHeap) Len() int {
	return len(*l)
}

func (l *listNodeHeap) Less(i, j int) bool {
	return (*l)[i].time > (*l)[j].time
}

func (l *listNodeHeap) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *listNodeHeap) Push(x interface{}) {
	*l = append(*l, x.(*listNode))
}

func (l *listNodeHeap) Pop() interface{} {
	n := l.Len()
	ret := (*l)[n-1]
	*l = (*l)[:n-1]
	return ret
}

type Twitter struct {
	time       int
	userTweets map[int]*listNode
	userFollow map[int]map[int]bool
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		userTweets: map[int]*listNode{},
		userFollow: map[int]map[int]bool{},
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userTweets[userId] = newListNode(tweetId, this.time, this.userTweets[userId])
	this.time++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := listNodeHeap{}
	// add followed users tweet
	for uid := range this.userFollow[userId] {
		if node := this.userTweets[uid]; node != nil {
			heap.Push(&h, node)
		}
	}

	// add user itself tweet
	if node := this.userTweets[userId]; node != nil {
		heap.Push(&h, node)
	}

	// return nil if no tweet
	if h.Len() == 0 {
		return nil
	}

	// get descending result
	i, res := 0, []int{}
	for ; i < 10 && h.Len() > 0; i++ {
		node := heap.Pop(&h).(*listNode)
		res = append(res, node.tweet)
		if node.next != nil {
			heap.Push(&h, node.next)
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}

	if set, ok := this.userFollow[followerId]; ok {
		set[followeeId] = true
	} else {
		set = map[int]bool{}
		set[followeeId] = true
		this.userFollow[followerId] = set
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	delete(this.userFollow[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
