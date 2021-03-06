// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package localcachelayer

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-server/v5/store/storetest"
	"github.com/mattermost/mattermost-server/v5/store/storetest/mocks"
)

func TestChannelStore(t *testing.T) {
	StoreTest(t, storetest.TestReactionStore)
}

func TestChannelStoreChannelMemberCountsCache(t *testing.T) {
	countResult := int64(10)

	t.Run("first call not cached, second cached and returning same data", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		count, err := cachedStore.Channel().GetMemberCount("id", true)
		require.Nil(t, err)
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
		count, err = cachedStore.Channel().GetMemberCount("id", true)
		require.Nil(t, err)
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
	})

	t.Run("first call not cached, second force not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
		cachedStore.Channel().GetMemberCount("id", false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 2)
	})

	t.Run("first call force not cached, second not cached, third cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetMemberCount("id", false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 2)
		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 2)
	})

	t.Run("first call with GetMemberCountFromCache not cached, second cached and returning same data", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		count := cachedStore.Channel().GetMemberCountFromCache("id")
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
		count = cachedStore.Channel().GetMemberCountFromCache("id")
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
	})

	t.Run("first call not cached, clear cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
		cachedStore.Channel().ClearCaches()
		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 2)
	})

	t.Run("first call not cached, invalidate cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 1)
		cachedStore.Channel().InvalidateMemberCount("id")
		cachedStore.Channel().GetMemberCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMemberCount", 2)
	})
}

func TestChannelStoreChannelPinnedPostsCountsCache(t *testing.T) {
	countResult := int64(10)

	t.Run("first call not cached, second cached and returning same data", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		count, err := cachedStore.Channel().GetPinnedPostCount("id", true)
		require.Nil(t, err)
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 1)
		count, err = cachedStore.Channel().GetPinnedPostCount("id", true)
		require.Nil(t, err)
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 1)
	})

	t.Run("first call not cached, second force not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 1)
		cachedStore.Channel().GetPinnedPostCount("id", false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 2)
	})

	t.Run("first call force not cached, second not cached, third cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetPinnedPostCount("id", false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 1)
		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 2)
		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 2)
	})

	t.Run("first call not cached, clear cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 1)
		cachedStore.Channel().ClearCaches()
		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 2)
	})

	t.Run("first call not cached, invalidate cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 1)
		cachedStore.Channel().InvalidatePinnedPostCount("id")
		cachedStore.Channel().GetPinnedPostCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetPinnedPostCount", 2)
	})
}

func TestChannelStoreGuestCountCache(t *testing.T) {
	countResult := int64(12)

	t.Run("first call not cached, second cached and returning same data", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		count, err := cachedStore.Channel().GetGuestCount("id", true)
		require.Nil(t, err)
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 1)
		count, err = cachedStore.Channel().GetGuestCount("id", true)
		require.Nil(t, err)
		assert.Equal(t, count, countResult)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 1)
	})

	t.Run("first call not cached, second force not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 1)
		cachedStore.Channel().GetGuestCount("id", false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 2)
	})

	t.Run("first call force not cached, second not cached, third cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetGuestCount("id", false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 1)
		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 2)
		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 2)
	})

	t.Run("first call not cached, clear cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 1)
		cachedStore.Channel().ClearCaches()
		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 2)
	})

	t.Run("first call not cached, invalidate cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 1)
		cachedStore.Channel().InvalidateGuestCount("id")
		cachedStore.Channel().GetGuestCount("id", true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetGuestCount", 2)
	})
}

func TestChannelStoreChannel(t *testing.T) {
	channelId := "channel1"
	fakeChannel := model.Channel{Id: channelId}
	t.Run("first call by id not cached, second cached and returning same data", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		channel, err := cachedStore.Channel().Get(channelId, true)
		require.Nil(t, err)
		assert.Equal(t, channel, &fakeChannel)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 1)
		channel, err = cachedStore.Channel().Get(channelId, true)
		require.Nil(t, err)
		assert.Equal(t, channel, &fakeChannel)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 1)
	})

	t.Run("first call not cached, second force no cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 1)
		cachedStore.Channel().Get(channelId, false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 2)
	})

	t.Run("first call force no cached, second not cached, third cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)
		cachedStore.Channel().Get(channelId, false)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 1)
		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 2)
		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 2)
	})

	t.Run("first call not cached, clear cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 1)
		cachedStore.Channel().ClearCaches()
		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 2)
	})

	t.Run("first call not cached, invalidate cache, second call not cached", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)
		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 1)
		cachedStore.Channel().InvalidateChannel(channelId)
		cachedStore.Channel().Get(channelId, true)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "Get", 2)
	})
}

func TestChannelStoreChannelMembersForUserCache(t *testing.T) {
	fakeChannelMembers := model.ChannelMembers([]model.ChannelMember{
		{
			UserId: "123",
		},
	})

	t.Run("first call not cached, second cached and returning same data", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		gotMembers, err := cachedStore.Channel().GetMembersForUser("teamId", "userId1")
		require.Nil(t, err)
		assert.Equal(t, &fakeChannelMembers, gotMembers)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 1)

		_, _ = cachedStore.Channel().GetMembersForUser("teamId", "userId1")
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 1)
	})

	t.Run("first call not cached, invalidate, and then not cached again", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		gotMembers, err := cachedStore.Channel().GetMembersForUser("teamId", "userId1")
		require.Nil(t, err)
		assert.Equal(t, &fakeChannelMembers, gotMembers)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 1)

		// invalidate
		cachedStore.Channel().(LocalCacheChannelStore).invalidateMembersForUser("userId1")

		_, _ = cachedStore.Channel().GetMembersForUser("teamId", "userId1")
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 2)
	})

	t.Run("cache multiple items, invalidate all, then nothing cached ", func(t *testing.T) {
		mockStore := getMockStore()
		mockCacheProvider := getMockCacheProvider()
		cachedStore := NewLocalCacheLayer(mockStore, nil, nil, mockCacheProvider)

		gotMembers, err := cachedStore.Channel().GetMembersForUser("teamId", "userId1")
		require.Nil(t, err)
		assert.Equal(t, &fakeChannelMembers, gotMembers)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 1)

		gotMembers, err = cachedStore.Channel().GetMembersForUser("teamId", "userId2")
		require.Nil(t, err)
		assert.Equal(t, &fakeChannelMembers, gotMembers)
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 2)

		// invalidate all
		cachedStore.Channel().(LocalCacheChannelStore).invalidateMembersForAllUsers()

		_, _ = cachedStore.Channel().GetMembersForUser("teamId", "userId1")
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 3)
		_, _ = cachedStore.Channel().GetMembersForUser("teamId", "userId2")
		mockStore.Channel().(*mocks.ChannelStore).AssertNumberOfCalls(t, "GetMembersForUser", 4)
	})
}
