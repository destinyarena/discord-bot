package handlers

func (h *handler) OnReady() func() {
	return func() {
		user, err := h.Bot.Client.CurrentUser().Get()
		if err != nil {
			h.Logger.Error("WE ARE FUCKED LOL")
			h.Logger.Fatal(err)
		}
		h.Logger.Infof("%s#%s", user.Username, user.Discriminator)
	}
}
