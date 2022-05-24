package mocks

//go:generate sh -c "mockgen -package mockquic -destination quic/stream.go github.com/For-ACGN/quic-bbr Stream && goimports -w quic/stream.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/session.go github.com/For-ACGN/quic-bbr Session && goimports -w quic/session.go"
//go:generate sh -c "../mockgen_internal.sh mocks sealer.go github.com/For-ACGN/quic-bbr/internal/handshake Sealer"
//go:generate sh -c "../mockgen_internal.sh mocks opener.go github.com/For-ACGN/quic-bbr/internal/handshake Opener"
//go:generate sh -c "../mockgen_internal.sh mocks crypto_setup.go github.com/For-ACGN/quic-bbr/internal/handshake CryptoSetup"
//go:generate sh -c "../mockgen_internal.sh mocks stream_flow_controller.go github.com/For-ACGN/quic-bbr/internal/flowcontrol StreamFlowController"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/sent_packet_handler.go github.com/For-ACGN/quic-bbr/internal/ackhandler SentPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/received_packet_handler.go github.com/For-ACGN/quic-bbr/internal/ackhandler ReceivedPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mocks congestion.go github.com/For-ACGN/quic-bbr/internal/congestion SendAlgorithmWithDebugInfos"
//go:generate sh -c "../mockgen_internal.sh mocks connection_flow_controller.go github.com/For-ACGN/quic-bbr/internal/flowcontrol ConnectionFlowController"
