package main

func main() {
	channel1()
	channelNoRecipient()
	channelNoRecipientV2()
	bufferedChannel()
	bufferedChannel2()
	directional1()
	/*channelReceiver := make(chan string, 1)
	channelReceiver <- "asdasd"
	receiverChannel(channelReceiver)
	fmt.Println(<-channelReceiver)*/

	selectAux()
	rangingChannels()
}
