type Ticket = {
  // code here
  from: string;
  to: string;
};

/**
 * Given an array of tickets:
   [
    { from: 'Astana', to: 'Bali' },
    { from: 'Dubai', to: 'Astana' },
    { from: 'Bali', to: 'Dublin' },
   ]
 
   You need to arrange them one after another to form a continuous route:
  [
    { from: 'Dubai', to: 'Astana' },
    { from: 'Astana', to: 'Bali' },
    { from: 'Bali', to: 'Dublin' }
  ]
 */

export const mockRoutes: Ticket[] = [
  {from: 'Astana', to: 'Bali'},
  {from: 'Dubai', to: 'Astana'},
  {from: 'Bali', to: 'Dublin'},
];

export const getRoute = (tickets = mockRoutes) => {
  // code here
  // connections map
  let connections: {[key: string]: string[]} = {};
  tickets.forEach(ticket => {
    if (connections.hasOwnProperty(ticket.from)) {
      connections[ticket.from] = [...connections[ticket.from], ticket.to];
    } else {
      connections[ticket.from] = [ticket.to];
    }
  });
  const fromArray = Object.keys(connections);
  const startingAirport = Object.values(connections)
    .map(toAirports => toAirports.find(airport => !fromArray.includes(airport)))
    .filter(v => v);
  console.log(connections);
  console.log('fromArray', fromArray);
  console.log('startingAirport', startingAirport);

  if (!startingAirport.length) {
    return "Error: We couldn't find staring airport";
  }
};
