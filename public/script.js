(function () {
  new Vue({
    el: "#app",
    data: {
      pickupInput: "",
      destinationInput: "",

      suggestions: [],
      suggestionType: "pickup",
      suggestionQuery: "",

      loading: false,

      result: {
        price: 0,
        distance: 0,
        duration: 0,
        fetched: false,
      },
    },

    methods: {
      async submit() {
        this.result = {
          price: 0,
          distance: 0,
          duration: 0,
          fetched: false,
        };
        if (this.pickupInput && this.destinationInput) {
          plotMap(this.pickupInput, this.destinationInput);

          const res = await getTripEstimate(
            this.pickupInput,
            this.destinationInput
          );

          this.result = {
            price: res.price,
            distance: res.distance,
            duration: res.duration,
            fetched: true,
          };
        }
      },

      clearPickupInput() {
        this.pickupInput = "";
        this.clearSuggestion();
      },

      clearDestinationInput() {
        this.destinationInput = "";
        this.clearSuggestion();
      },

      fetchPickupPrediction(e) {
        const input = e.target.value;
        this.fetchPrediction(input, "pickup");
      },

      fetchDestinationPrediction(e) {
        const input = e.target.value;
        this.fetchPrediction(input, "destination");
      },

      clearSuggestion() {
        this.suggestions = [];
        this.suggestionType = "pickup";
        this.suggestionQuery = "";
      },

      populatePickupInput(value) {
        this.pickupInput = value;
        this.clearSuggestion();
      },

      populateDestinationInput(value) {
        this.destinationInput = value;
        this.clearSuggestion();
      },

      fetchPrediction(input, type) {
        this.suggestionQuery = input;
        this.suggestionType = type;

        if (input === "") {
          this.clearSuggestion();
          return;
        }

        setTimeout(async () => {
          if (input !== this.suggestionQuery) {
            return;
          }

          this.suggestions = await getAddressSuggestion(input);
        }, 500);
      },
    },
  });
})();
