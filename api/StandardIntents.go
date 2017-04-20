package api

// The assistant fires the main intent for the root level call (for example "Ok Google, Talk To <Action>").
const MAIN_INTENT = "assistant.intent.action.MAIN"

// The assistant fires the text intent when you the action ask for a response.
const TEXT_INTENT = "assistant.intent.action.TEXT"

// The assistant fires the permission intent when an action invokes askForPermission
const PERMISSION_INTENT = "assistant.intent.action.PERMISSION"
