import * as core from "@actions/core";

const run = async (): Promise<void> => {
  try {
    console.log("hello world");
  } catch (e) {
    if (e instanceof Error) core.setFailed(e.message);
  }
};

run();
