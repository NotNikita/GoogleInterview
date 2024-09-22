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

export function getRoute(tickets = mockRoutes) {
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
  const endingAirport = Object.values(connections)
    .map(toAirports => toAirports.find(airport => !fromArray.includes(airport)))
    .filter(v => v);

  if (!endingAirport || !endingAirport.length) {
    return "Error: We couldn't find staring airport";
  }
  let resultRoute = [endingAirport[0]];
  let currAirport: string | undefined = endingAirport[0];
  // we are starting from the end (reverse BFS?)
  while (currAirport && typeof currAirport === 'string') {
    const prevAirportKV = Object.entries(connections).find(([key, value]) => {
      return value.includes(currAirport as string) ? key : false;
    });
    if (prevAirportKV) {
      const [prevAirportName, _] = prevAirportKV;

      resultRoute.push(prevAirportName);
      currAirport = prevAirportName;
    } else {
      // We reached the end
      currAirport = undefined;
    }
  }

  return resultRoute.reverse();
}

export function getRouteFaster(tickets = mockRoutes) {
  // const destinationAirports = tickets.map(t => t.to);
  const destinationAirports = new Set(tickets.map(t => t.to));
  // the start airport is the one, that was never on 'to' key. Opposit goes to end airport
  const startAirport = tickets.find(ticket => !destinationAirports.has(ticket.from))?.from;

  // Create a map for quick lookup of the next airport
  const airportMap = new Map(tickets.map(ticket => [ticket.from, ticket.to]));

  let result = [startAirport];
  let currentAirport: string | undefined = startAirport;

  while (currentAirport) {
    const someNextAirport = airportMap.get(currentAirport);

    if (someNextAirport) {
      currentAirport = someNextAirport;
      result.push(someNextAirport);
    } else {
      // stop cycle;
      currentAirport = someNextAirport;
    }
  }

  return result;
}
