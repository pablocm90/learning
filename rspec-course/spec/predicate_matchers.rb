RSpec.describe 'predicate methods and predicate matchers' do
  it 'can be tested with predicate matchers ' do
    result = 16 / 2
    expect(result).to be_even
  end
end
