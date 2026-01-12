#!/usr/bin/env bun
import { Command } from 'commander';
import { initCommand } from './commands/init';

const program = new Command();

program.name('ai').description('AI-driven repository setup and management tool').version('0.1.0');

program.addCommand(initCommand);

program.parse();
