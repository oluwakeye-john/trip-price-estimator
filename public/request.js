const getAddressSuggestion = (input) => {
  const query = `
        query{
            addressSuggestion(input: "${input}"){
                predictions{
                    description
                }
            }
        }
    `;

  return fetch("/query", {
    method: "post",
    body: JSON.stringify({ query }),
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((res) => res.json())
    .then((res) => {
      return res.data.addressSuggestion.predictions;
    })
    .catch(() => {
      return [];
    });
};

const getTripEstimate = (origin, destination) => {
  const query = `
        query{
          tripEstimate(origin: "${origin}", destination: "${destination}"){
            price
            distance
            duration
          }
        }
    `;

  return fetch("/query", {
    method: "post",
    body: JSON.stringify({ query }),
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((res) => res.json())
    .then((res) => {
      return res.data.tripEstimate;
    })
    .catch(() => {
      return [];
    });
};
