const TESTS_GLOB = ["**/__tests__/**", "**/*.test.*", "**/*.spec.*"];

module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-hooks/recommended",
    "prettier",
  ],
  ignorePatterns: ["dist", ".eslintrc.cjs"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "tsconfig.json",
    sourceType: "module",
  },
  plugins: [
    "react-refresh",
    "import",
    "react",
    "unused-imports",
    "no-only-tests",
    "only-warn",
  ],
  rules: {
    "react-refresh/only-export-components": [
      "warn",
      { allowConstantExport: true },
    ],
  },
  settings: {
    react: {
      version: "detect",
    },
  },
  root: true,
  overrides: [
    {
      files: TESTS_GLOB,
    },
    {
      files: ["*.js", "*.jsx"],
      extends: ["plugin:@typescript-eslint/disable-type-checked"],
    },
  ],

  rules: {
    "@typescript-eslint/naming-convention": [
      "warn",
      {
        selector: "enumMember",
        format: ["PascalCase"],
      },
      {
        selector: "enum",
        format: ["PascalCase"],
      },
    ],
    "@typescript-eslint/no-shadow": "warn",
    "no-unused-expressions": "off",
    "@typescript-eslint/no-unused-expressions": "warn",
    "@typescript-eslint/no-redeclare": "warn",
    "@typescript-eslint/unbound-method": "off", // Complains about i18n formatMessage method
    "@typescript-eslint/no-unnecessary-condition": "on",
    // ^ This is a very good rule and it would be the best to turn it on
    // However, it will auto-fix all cases where we, for example, use an unnecessary optional chain
    // This is a runtime change and could be unsafe if our types are wrong, for this reason it should be done in another MR
    "@typescript-eslint/no-confusing-void-expression": [
      "warn",
      { ignoreArrowShorthand: true },
    ],

    "unused-imports/no-unused-imports": "error",
    "import/no-extraneous-dependencies": [
      "warn",
      {
        devDependencies: TESTS_GLOB,
      },
    ],
    "import/no-internal-modules": "off",
    "import/no-relative-parent-imports": "error",
    "no-restricted-imports": [
      "error",
      {
        patterns: [
          {
            group: ["../../*"],
            message: "Usage of relative parent imports is not allowed.",
          },
        ],
      },
    ],
    "import/order": [
      "error",
      {
        groups: [
          "builtin",
          "external",
          "internal",
          ["parent", "sibling", "index"],
          "object",
        ],
        pathGroups: [
          {
            pattern: "constants/**",
            group: "internal",
          },
          {
            pattern: "ui/**",
            group: "internal",
            position: "after",
          },
          {
            pattern: "containers/**",
            group: "internal",
            position: "after",
          },
          {
            pattern: "components/**",
            group: "internal",
            position: "after",
          },
        ],
        pathGroupsExcludedImportTypes: ["react", "constants"],
        distinctGroup: false,
        "newlines-between": "always-and-inside-groups",
      },
    ],
    "import/no-cycle": "error",

    eqeqeq: ["error", "smart"],
    "guard-for-in": "error",
    "no-bitwise": "error",
    "no-caller": "error",
    "no-console": ["error", { allow: ["warn", "error"] }],
    "no-duplicate-imports": "error",
    "no-empty-function": "error",
    "no-eval": "error",
    "no-extra-bind": "error",
    "no-new-func": "error",
    "no-new-wrappers": "error",
    "no-sequences": "error",
    "no-template-curly-in-string": "error",
    "no-throw-literal": "error",
    "no-undef-init": "error",
    "object-shorthand": "error",
    "prefer-object-spread": "error",
    radix: "off",

    "react/jsx-boolean-value": "error",
    "react/self-closing-comp": "error",
    "react/react-in-jsx-scope": "off",
    "react/prop-types": [1, { ignore: ["children"] }],

    "no-only-tests/no-only-tests": "error",
  },
};
