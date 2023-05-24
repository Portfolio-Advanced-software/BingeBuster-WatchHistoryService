package handlers

import (
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
)

type HistoryServiceServer struct {
	historypb.UnimplementedHistoryServiceServer
}
