RSpec.describe ' the include matcher ' do
  describe 'hot chocolate' do
    it 'checks for substrings' do
      expect(subject).to include('hot')
      expect(subject).to include('choc')
      expect(subject).to include('late')
    end
    it { is_expected.to include('choc') }
  end

  describe [1, 2, 3, 4] do
    it 'checks for sub arrays' do
      expect(subject).to include(1, 2)
      expect(subject).to include(2, 4)
    end
    it { is_expected.to include(2, 3) }
  end

  describe({ a: 1, b: 2, c: 3 }) do
    it 'checks for sub hashes' do
      expect(subject).to include({ a: 1, b: 2 })
    end
    it { is_expected.to include({ b: 2 }) }
    it { is_expected.to include({ a: 1, c: 3 }) }
    it { is_expected.to include(:a, :c) }
  end
end
