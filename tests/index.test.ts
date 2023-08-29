import { hello } from "../src/util";

describe('Hello world', () => {
    it('hello -> world', () => {
        expect(hello()).toBe('world');
    });
});