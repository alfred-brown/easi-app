import React from 'react';
import { act } from 'react-dom/test-utils';
import axios from 'axios';
import { FlagProvider, FlagToggle, useFlags } from 'contexts/flagContext';
import { mount } from 'enzyme';

jest.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

beforeEach(() => {
  mockedAxios.get.mockClear();
});

const FlagPrinter = () => {
  const flags = useFlags();
  const stringified = JSON.stringify(flags, Object.keys(flags).sort());
  return <p>{stringified}</p>;
};

it('falls back to default values without a provider', () => {
  const component = mount(
    <FlagToggle name="sandbox">
      <p>Sandbox</p>
    </FlagToggle>
  );

  expect(component.find('p').exists()).toEqual(false);
});

it('loads flags into the provider', async () => {
  const getMock = mockedAxios.get.mockResolvedValue({
    data: {
      sandbox: true,
      taskListLite: true
    }
  });

  const wrapper = mount(
    <FlagProvider>
      <FlagPrinter />
    </FlagProvider>
  );
  await act(async () => {
    await getMock;
  });

  wrapper.update();

  const printer = wrapper.find(FlagPrinter);
  expect(printer.text()).toEqual(`{"sandbox":true,"taskListLite":true}`);
  expect(mockedAxios.get.mock.calls).toEqual([
    ['http://localhost:8080/api/v1/flags']
  ]);
});

it('uses the defaults when flags fail to load', async () => {
  const getMock = mockedAxios.get.mockRejectedValue(
    new Error('could not load')
  );

  const wrapper = mount(
    <FlagProvider>
      <FlagPrinter />
    </FlagProvider>
  );
  await act(async () => {
    await getMock;
  });

  wrapper.update();

  const printer = wrapper.find(FlagPrinter);
  expect(printer.text()).toEqual(`{"sandbox":false,"taskListLite":false}`);
  expect(mockedAxios.get.mock.calls).toEqual([
    ['http://localhost:8080/api/v1/flags']
  ]);
});
