import { Attribute, DeliverTxResponse, Event } from "@cosmjs/stargate/build";

export type GameCreatedEvent = Event;

export const getCreateGameEvent = (
  res: DeliverTxResponse
): GameCreatedEvent | undefined =>
  res.events?.find(
    (event: Event) => event.type === "lb.rps.v1.EventCreateGame"
  );

export const getCreatedGameId = (
  createdGameEvent: GameCreatedEvent
): number | undefined => {
  const gameNum = createdGameEvent.attributes.find(
    (attribute: Attribute) => attribute.key == "game_number"
  )!.value;
  // attributes values are wrapped with additional doble quotes (")
  // check if the issue is on the Cosmos-SDK side
  const sanitizedGameNum = gameNum.replace('"', "");
  return gameNum ? parseInt(sanitizedGameNum, 10) : undefined;
};

export type MakeMoveEvent = Event;

export const getMakeMoveEvent = (
  res: DeliverTxResponse
): MakeMoveEvent | undefined =>
  res.events?.find((event: Event) => event.type === "lb.rps.v1.EventMakeMove");

export const getMove = (makeMoveEvent: MakeMoveEvent): string => {
  const move = makeMoveEvent.attributes.find(
    (attribute: Attribute) => attribute.key == "move"
  )!.value;
  // attributes values are wrapped with additional doble quotes (")
  // check if the issue is on the Cosmos-SDK side
  return move.replace('"', "");
};

export type RevealMoveEvent = Event;

export const getRevealMoveEvent = (
  res: DeliverTxResponse
): RevealMoveEvent | undefined =>
  res.events?.find(
    (event: Event) => event.type === "lb.rps.v1.EventRevealMove"
  );

export const getRevealedMove = (revealMoveEvent: RevealMoveEvent): string => {
  const move = revealMoveEvent.attributes.find(
    (attribute: Attribute) => attribute.key == "revealed_move"
  )!.value;

  // attributes values are wrapped with additional doble quotes (")
  // check if the issue is on the Cosmos-SDK side
  return move.replace('"', "");
};

export type GameEndedEvent = Event;

export const getGameEndedEvent = (
  res: DeliverTxResponse
): GameEndedEvent | undefined =>
  res.events?.find((event: Event) => event.type === "lb.rps.v1.EventEndGame");

export const getEndedGameId = (
  gameEndedEvent: GameEndedEvent
): number | undefined => {
  const gameNum = gameEndedEvent.attributes.find(
    (attribute: Attribute) => attribute.key == "game_number"
  )!.value;
  // attributes values are wrapped with additional doble quotes (")
  // check if the issue is on the Cosmos-SDK side
  const sanitizedGameNum = gameNum.replace('"', "");
  return gameNum ? parseInt(sanitizedGameNum, 10) : undefined;
};

export const getEndedGameStatus = (gameEndedEvent: GameEndedEvent): string => {
  const status = gameEndedEvent.attributes.find(
    (attribute: Attribute) => attribute.key == "status"
  )!.value;

  // attributes values are wrapped with additional doble quotes (")
  // check if the issue is on the Cosmos-SDK side
  return status.replace('"', "");
};